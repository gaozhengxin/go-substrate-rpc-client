package types

import "github.com/centrifuge/go-substrate-rpc-client/v4/scale"

type CallOf Call

type DispatchTime struct {
	IsAt    bool
	AsAt    BlockNumber
	IsAfter bool
	AsAfter BlockNumber
}

func (m *DispatchTime) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		m.IsAt = true
		err = decoder.Decode(&m.AsAt)
	} else if b == 1 {
		m.IsAfter = true
		err = decoder.Decode(&m.AsAfter)
	}

	if err != nil {
		return err
	}

	return nil
}

func (m DispatchTime) Encode(encoder scale.Encoder) error {
	var err1, err2 error
	if m.IsAt {
		err1 = encoder.PushByte(0)
		err2 = encoder.Encode(&m.AsAt)
	} else if m.IsAfter {
		err1 = encoder.PushByte(1)
		err2 = encoder.Encode(&m.AsAfter)
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}

type ScheduleTaskIndex U32

type DelayedOrigin struct {
	Delay  BlockNumber
	Origin PalletsOrigin
}

type AuthorityOrigin DelayedOrigin

type StorageValue []U8

type GraduallyUpdate struct {
	Key         StorageValue
	TargetValue StorageValue
	PerBlock    StorageValue
}

type StorageKeyBytes []U8

type StorageValueBytes []U8

type RpcDataProviderId Text

type DataProviderId U8

type TimestampedValue struct {
	Value     OracleValue
	Timestamp Moment
}

type TimestampedValueOf TimestampedValue

type OrderedSet []AccountID

type OrmlAccountData struct {
	Free     Balance
	Frozen   Balance
	Reserved Balance
}

type OrmlBalanceLock struct {
	Amount Balance
	Id     LockIdentifier
}

type LockIdentifier U8

type AuctionInfo struct {
	Bid   OptionBid
	Start BlockNumber
	End   OptionBlockNumber
}

type OptionBid struct {
	option
	value Bid
}

func (o OptionBid) Unwrap() (ok bool, value Bid) {
	return o.hasValue, o.value
}

func (o OptionBid) Encode(encoder scale.Encoder) error {
	return encoder.EncodeOption(o.hasValue, o.value)
}

func (o *OptionBid) Decode(decoder scale.Decoder) error {
	return decoder.DecodeOption(&o.hasValue, &o.value)
}

type Bid struct {
	AccountID
	Balance
}

type OptionBlockNumber struct {
	option
	value BlockNumber
}

func (o OptionBlockNumber) Unwrap() (ok bool, value BlockNumber) {
	return o.hasValue, o.value
}

func (o OptionBlockNumber) Encode(encoder scale.Encoder) error {
	return encoder.EncodeOption(o.hasValue, o.value)
}

func (o *OptionBlockNumber) Decode(decoder scale.Decoder) error {
	return decoder.DecodeOption(&o.hasValue, &o.value)
}

type DelayedDispatchTime struct {
	IsAt    bool
	AsAt    BlockNumber
	IsAfter bool
	AsAfter BlockNumber
}

func (m *DelayedDispatchTime) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		m.IsAt = true
		err = decoder.Decode(&m.AsAt)
	} else if b == 1 {
		m.IsAfter = true
		err = decoder.Decode(&m.AsAfter)
	}

	if err != nil {
		return err
	}

	return nil
}

func (m DelayedDispatchTime) Encode(encoder scale.Encoder) error {
	var err1, err2 error
	if m.IsAt {
		err1 = encoder.PushByte(0)
		err2 = encoder.Encode(&m.AsAt)
	} else if m.IsAfter {
		err1 = encoder.PushByte(1)
		err2 = encoder.Encode(&m.AsAfter)
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}

type DispatchId U32

type Price U128

type OrmlVestingSchedule struct {
	Start       BlockNumber
	Peroid      BlockNumber
	PeriodCount U32
	PerPeriod   CompactBalance
}

type VestingScheduleOf OrmlVestingSchedule

type OrmlCurrencyId U8

type PoolInfo struct{}

type CompactBalance Balance

type BalanceOf Balance

type Balance U256

type PoolInfoV0 struct {
	TotalShares           Share
	TotalRewards          CompactBalance
	TotalWithdrawnRewards CompactBalance
}

type Share U128

type OracleValue Price

type PalletBalanceOf Balance

type EstimateResourcesResponse struct {
	Gas       U256
	Storage   I32
	WeightFee U256
}

type EvmAccountInfo struct {
	Nonce            Index
	ContractInfo     OptionEvmContractInfo
	DeveloperDeposit OptionBalance
}

type Index U256

type OptionEvmContractInfo struct {
	option
	value EvmContractInfo
}

func (o OptionEvmContractInfo) Unwrap() (ok bool, value EvmContractInfo) {
	return o.hasValue, o.value
}

func (o OptionEvmContractInfo) Encode(encoder scale.Encoder) error {
	return encoder.EncodeOption(o.hasValue, o.value)
}

func (o *OptionEvmContractInfo) Decode(decoder scale.Decoder) error {
	return decoder.DecodeOption(&o.hasValue, &o.value)
}

type OptionBalance struct {
	option
	value Balance
}

func (o OptionBalance) Encode(encoder scale.Encoder) error {
	return encoder.EncodeOption(o.hasValue, o.value)
}

func (o *OptionBalance) Decode(decoder scale.Decoder) error {
	return decoder.DecodeOption(&o.hasValue, &o.value)
}

func (o OptionBalance) Unwrap() (ok bool, value Balance) {
	return o.hasValue, o.value
}

type CodeInfo struct {
	CodeSize U32
	RefCount U32
}

type EvmContractInfo struct {
	CodeHash   H256
	Maintainer H160
	Deployed   Bool
}

type EvmAddress H160

type CallRequest struct {
	From         OptionH160
	To           OptionH160
	GasLimit     OptionU32
	StorageLimit OptionU32
	Value        OptionU128
	Data         OptionBytes
}

type CommitmentOf struct {
	Duration  LockDuration
	Amount    BalanceOf
	Candidate AccountID
}

type Era struct {
	Index U32
	Start BlockNumber
}

type LockDuration string

type PoolId struct {
	IsLoans        bool
	AsLoans        CurrencyId
	IsDexIncentive bool
	AsDexIncentive CurrencyId
	IsDexSaving    bool
	AsDexSaving    CurrencyId
	IsHoma         bool
	AsHoma         Null
}

func (m *PoolId) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		m.IsLoans = true
		err = decoder.Decode(&m.AsLoans)
	} else if b == 1 {
		m.IsDexIncentive = true
		err = decoder.Decode(&m.AsDexIncentive)
	} else if b == 2 {
		m.IsDexSaving = true
		err = decoder.Decode(&m.AsDexSaving)
	} else if b == 3 {
		m.IsHoma = true
		err = decoder.Decode(&m.AsHoma)
	}

	if err != nil {
		return err
	}

	return nil
}

func (m PoolId) Encode(encoder scale.Encoder) error {
	var err1, err2 error
	if m.IsLoans {
		err1 = encoder.PushByte(0)
		err2 = encoder.Encode(&m.AsLoans)
	} else if m.IsDexIncentive {
		err1 = encoder.PushByte(1)
		err2 = encoder.Encode(&m.AsDexIncentive)
	} else if m.IsDexSaving {
		err1 = encoder.PushByte(2)
		err2 = encoder.Encode(&m.AsDexSaving)
	} else if m.IsHoma {
		err1 = encoder.PushByte(3)
		err2 = encoder.Encode(&m.AsHoma)
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}

type Amount I128

type AmountOf Amount

type AuctionId U32

type AuctionOf AuctionId

type TokenSymbol U8

type CurrencyId struct {
	IsToken    bool
	AsToken    TokenSymbol
	IsDEXShare bool
	AsDEXShare TypeDEXShare
	IsERC20    bool
	AsERC20    EvmAddress
}

func (m *CurrencyId) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		m.IsToken = true
		err = decoder.Decode(&m.AsToken)
	} else if b == 1 {
		m.IsDEXShare = true
		err = decoder.Decode(&m.AsDEXShare)
	} else if b == 2 {
		m.IsERC20 = true
		err = decoder.Decode(&m.AsERC20)
	}

	if err != nil {
		return err
	}

	return nil
}

func (m CurrencyId) Encode(encoder scale.Encoder) error {
	var err1, err2 error
	if m.IsToken {
		err1 = encoder.PushByte(0)
		err2 = encoder.Encode(&m.AsToken)
	} else if m.IsDEXShare {
		err1 = encoder.PushByte(1)
		err2 = encoder.Encode(&m.AsDEXShare)
	} else if m.IsERC20 {
		err1 = encoder.PushByte(2)
		err2 = encoder.Encode(&m.AsERC20)
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}

type TypeDEXShare struct {
	Token1 TokenSymbol
	Token2 TokenSymbol
}

type CurrencyIdOf CurrencyId

type AuthoritysOriginId string

func (m *AuthoritysOriginId) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		err = decoder.Decode(&m)
	}

	if err != nil {
		return err
	}

	return nil
}

func (m AuthoritysOriginId) Encode(encoder scale.Encoder) error {
	var err1, err2 error
	if m == "Root" {
		err1 = encoder.PushByte(0)
		err2 = encoder.Encode(&m)
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}

type TradingPair struct {
	Currency1 CurrencyId
	Currency2 CurrencyId
}

type OracleKey CurrencyId

type AsOriginId AuthoritysOriginId

type ExchangeRate U128

type SystemOrigin interface{}

type CollectiveOrigin interface{}

type PalletsOrigin struct {
	IsSystem                       bool
	AsSystem                       SystemOrigin
	IsTimestamp                    bool
	AsTimestamp                    Null
	IsRandomnessCollectiveFlip     bool
	AsRandomnessCollectiveFlip     Null
	IsBalances                     bool
	AsBalances                     Null
	IsAccounts                     bool
	AsAccounts                     Null
	IsCurrencies                   bool
	AsCurrencies                   Null
	IsTokens                       bool
	AsTokens                       Null
	IsVesting                      bool
	AsVesting                      Null
	IsAcalaTreasury                bool
	AsAcalaTreasury                Null
	IsUtility                      bool
	AsUtility                      Null
	IsMultisig                     bool
	AsMultisig                     Null
	IsRecovery                     bool
	AsRecovery                     Null
	IsProxy                        bool
	AsProxy                        Null
	IsScheduler                    bool
	AsScheduler                    Null
	IsIndices                      bool
	AsIndices                      Null
	IsGraduallyUpdate              bool
	AsGraduallyUpdate              Null
	IsAuthorship                   bool
	AsAuthorship                   Null
	IsBabe                         bool
	AsBabe                         Null
	IsGrandpa                      bool
	AsGrandpa                      Null
	IsStaking                      bool
	AsStaking                      Null
	IsSession                      bool
	AsSession                      Null
	IsHistorical                   bool
	AsHistorical                   Null
	IsGeneralCouncil               bool
	AsGeneralCouncil               CollectiveOrigin
	IsGeneralCouncilMembership     bool
	AsGeneralCouncilMembership     Null
	IsHonzonCouncil                bool
	AsHonzonCouncil                CollectiveOrigin
	IsHonzonCouncilMembership      bool
	AsHonzonCouncilMembership      Null
	IsHomaCouncil                  bool
	AsHomaCouncil                  CollectiveOrigin
	IsHomaCouncilMembership        bool
	AsHomaCouncilMembership        Null
	IsTechnicalCommittee           bool
	AsTechnicalCommittee           Null
	IsTechnicalCommitteeMembership bool
	AsTechnicalCommitteeMembership Null
	IsAuthority                    bool
	//AsAuthority                    DelayedOrigin
	AsAuthority               Null
	IsElectionsPhragmen       bool
	AsElectionsPhragmen       Null
	IsAcalaOracle             bool
	AsAcalaOracle             Null
	IsBandOracle              bool
	AsBandOracle              Null
	IsOperatorMembershipAcala bool
	AsOperatorMembershipAcala Null
	IsOperatorMembershipBand  bool
	AsOperatorMembershipBand  Null
	IsAuction                 bool
	AsAuction                 Null
	IsRewards                 bool
	AsRewards                 Null
	IsOrmlNFT                 bool
	AsOrmlNFT                 Null
	IsPrices                  bool
	AsPrices                  Null
	IsDex                     bool
	AsDex                     Null
	IsAuctionManager          bool
	AsAuctionManager          Null
	IsLoans                   bool
	AsLoans                   Null
	IsHonzon                  bool
	AsHonzon                  Null
	IsCdpTreasury             bool
	AsCdpTreasury             Null
	IsCdpEngine               bool
	AsCdpEngine               Null
	IsEmergencyShutdown       bool
	AsEmergencyShutdown       Null
	IsHoma                    bool
	AsHoma                    Null
	IsNomineesElection        bool
	AsNomineesElection        Null
	IsStakingPool             bool
	AsStakingPool             Null
	IsPolkadotBridge          bool
	AsPolkadotBridge          Null
	IsIncentives              bool
	AsIncentives              Null
	IsAirDrop                 bool
	AsAirDrop                 Null
	IsNFT                     bool
	AsNFT                     Null
	IsRenVmBridge             bool
	AsRenVmBridge             Null
	IsContracts               bool
	AsContracts               Null
	IsEVM                     bool
	AsEVM                     Null
	IsSudo                    bool
	AsSudo                    Null
	IsTransactionPayment      bool
	AsTransactionPayment      Null
}

func (m *PalletsOrigin) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		m.IsSystem = true
		err = decoder.Decode(&m.AsSystem)
	} else if b == 1 {
		m.IsTimestamp = true
		err = decoder.Decode(&m.AsTimestamp)
	} else if b == 2 {
		m.IsRandomnessCollectiveFlip = true
		err = decoder.Decode(&m.AsRandomnessCollectiveFlip)
	} else if b == 3 {
		m.IsBalances = true
		err = decoder.Decode(&m.AsBalances)
	} else if b == 4 {
		m.IsAccounts = true
		err = decoder.Decode(&m.AsAccounts)
	} else if b == 5 {
		m.IsCurrencies = true
		err = decoder.Decode(&m.AsCurrencies)
	} else if b == 6 {
		m.IsTokens = true
		err = decoder.Decode(&m.AsTokens)
	} else if b == 7 {
		m.IsVesting = true
		err = decoder.Decode(&m.AsVesting)
	} else if b == 8 {
		m.IsAcalaTreasury = true
		err = decoder.Decode(&m.AsAcalaTreasury)
	} else if b == 9 {
		m.IsUtility = true
		err = decoder.Decode(&m.AsUtility)
	} else if b == 10 {
		m.IsMultisig = true
		err = decoder.Decode(&m.AsMultisig)
	} else if b == 11 {
		m.IsRecovery = true
		err = decoder.Decode(&m.AsRecovery)
	} else if b == 12 {
		m.IsProxy = true
		err = decoder.Decode(&m.AsProxy)
	} else if b == 13 {
		m.IsScheduler = true
		err = decoder.Decode(&m.AsScheduler)
	} else if b == 14 {
		m.IsIndices = true
		err = decoder.Decode(&m.AsIndices)
	} else if b == 15 {
		m.IsGraduallyUpdate = true
		err = decoder.Decode(&m.AsGraduallyUpdate)
	} else if b == 16 {
		m.IsAuthorship = true
		err = decoder.Decode(&m.AsAuthorship)
	} else if b == 17 {
		m.IsBabe = true
		err = decoder.Decode(&m.AsBabe)
	} else if b == 18 {
		m.IsGrandpa = true
		err = decoder.Decode(&m.AsGrandpa)
	} else if b == 19 {
		m.IsStaking = true
		err = decoder.Decode(&m.AsStaking)
	} else if b == 20 {
		m.IsSession = true
		err = decoder.Decode(&m.AsSession)
	} else if b == 21 {
		m.IsHistorical = true
		err = decoder.Decode(&m.AsHistorical)
	} else if b == 22 {
		m.IsGeneralCouncil = true
		err = decoder.Decode(&m.AsGeneralCouncil)
	} else if b == 23 {
		m.IsGeneralCouncilMembership = true
		err = decoder.Decode(&m.AsGeneralCouncilMembership)
	} else if b == 24 {
		m.IsHonzonCouncil = true
		err = decoder.Decode(&m.AsHonzonCouncil)
	} else if b == 25 {
		m.IsHonzonCouncilMembership = true
		err = decoder.Decode(&m.AsHonzonCouncilMembership)
	} else if b == 26 {
		m.IsHomaCouncil = true
		err = decoder.Decode(&m.AsHomaCouncil)
	} else if b == 27 {
		m.IsHomaCouncilMembership = true
		err = decoder.Decode(&m.AsHomaCouncilMembership)
	} else if b == 28 {
		m.IsTechnicalCommittee = true
		err = decoder.Decode(&m.AsTechnicalCommittee)
	} else if b == 29 {
		m.IsTechnicalCommitteeMembership = true
		err = decoder.Decode(&m.AsTechnicalCommitteeMembership)
	} else if b == 30 {
		m.IsAuthority = true
		err = decoder.Decode(&m.AsAuthority)
	} else if b == 31 {
		m.IsElectionsPhragmen = true
		err = decoder.Decode(&m.AsElectionsPhragmen)
	} else if b == 32 {
		m.IsAcalaOracle = true
		err = decoder.Decode(&m.AsAcalaOracle)
	} else if b == 33 {
		m.IsBandOracle = true
		err = decoder.Decode(&m.AsBandOracle)
	} else if b == 34 {
		m.IsOperatorMembershipAcala = true
		err = decoder.Decode(&m.AsOperatorMembershipAcala)
	} else if b == 35 {
		m.IsOperatorMembershipBand = true
		err = decoder.Decode(&m.AsOperatorMembershipBand)
	} else if b == 36 {
		m.IsAuction = true
		err = decoder.Decode(&m.AsAuction)
	} else if b == 37 {
		m.IsRewards = true
		err = decoder.Decode(&m.AsRewards)
	} else if b == 38 {
		m.IsOrmlNFT = true
		err = decoder.Decode(&m.AsOrmlNFT)
	} else if b == 39 {
		m.IsPrices = true
		err = decoder.Decode(&m.AsPrices)
	} else if b == 40 {
		m.IsDex = true
		err = decoder.Decode(&m.AsDex)
	} else if b == 41 {
		m.IsAuctionManager = true
		err = decoder.Decode(&m.AsAuctionManager)
	} else if b == 42 {
		m.IsLoans = true
		err = decoder.Decode(&m.AsLoans)
	} else if b == 43 {
		m.IsHonzon = true
		err = decoder.Decode(&m.AsHonzon)
	} else if b == 44 {
		m.IsCdpTreasury = true
		err = decoder.Decode(&m.AsCdpTreasury)
	} else if b == 45 {
		m.IsCdpEngine = true
		err = decoder.Decode(&m.AsCdpEngine)
	} else if b == 46 {
		m.IsEmergencyShutdown = true
		err = decoder.Decode(&m.AsEmergencyShutdown)
	} else if b == 47 {
		m.IsHoma = true
		err = decoder.Decode(&m.AsHoma)
	} else if b == 48 {
		m.IsNomineesElection = true
		err = decoder.Decode(&m.AsNomineesElection)
	} else if b == 49 {
		m.IsStakingPool = true
		err = decoder.Decode(&m.AsStakingPool)
	} else if b == 50 {
		m.IsPolkadotBridge = true
		err = decoder.Decode(&m.AsPolkadotBridge)
	} else if b == 51 {
		m.IsIncentives = true
		err = decoder.Decode(&m.AsIncentives)
	} else if b == 52 {
		m.IsAirDrop = true
		err = decoder.Decode(&m.AsAirDrop)
	} else if b == 53 {
		m.IsNFT = true
		err = decoder.Decode(&m.AsNFT)
	} else if b == 54 {
		m.IsRenVmBridge = true
		err = decoder.Decode(&m.AsRenVmBridge)
	} else if b == 55 {
		m.IsContracts = true
		err = decoder.Decode(&m.AsContracts)
	} else if b == 56 {
		m.IsEVM = true
		err = decoder.Decode(&m.AsEVM)
	} else if b == 57 {
		m.IsSudo = true
		err = decoder.Decode(&m.AsSudo)
	} else if b == 58 {
		m.IsTransactionPayment = true
		err = decoder.Decode(&m.AsTransactionPayment)
	}

	if err != nil {
		return err
	}

	return nil
}

func (m PalletsOrigin) Encode(encoder scale.Encoder) error {
	var err1, err2 error
	if m.IsSystem {
		err1 = encoder.PushByte(0)
		err2 = encoder.Encode(&m.AsSystem)
	} else if m.IsTimestamp {
		err1 = encoder.PushByte(1)
		err2 = encoder.Encode(&m.AsTimestamp)
	} else if m.IsRandomnessCollectiveFlip {
		err1 = encoder.PushByte(2)
		err2 = encoder.Encode(&m.AsRandomnessCollectiveFlip)
	} else if m.IsBalances {
		err1 = encoder.PushByte(3)
		err2 = encoder.Encode(&m.AsBalances)
	} else if m.IsAccounts {
		err1 = encoder.PushByte(4)
		err2 = encoder.Encode(&m.AsAccounts)
	} else if m.IsCurrencies {
		err1 = encoder.PushByte(5)
		err2 = encoder.Encode(&m.AsCurrencies)
	} else if m.IsTokens {
		err1 = encoder.PushByte(6)
		err2 = encoder.Encode(&m.AsTokens)
	} else if m.IsVesting {
		err1 = encoder.PushByte(7)
		err2 = encoder.Encode(&m.AsVesting)
	} else if m.IsAcalaTreasury {
		err1 = encoder.PushByte(8)
		err2 = encoder.Encode(&m.AsAcalaTreasury)
	} else if m.IsUtility {
		err1 = encoder.PushByte(9)
		err2 = encoder.Encode(&m.AsUtility)
	} else if m.IsMultisig {
		err1 = encoder.PushByte(10)
		err2 = encoder.Encode(&m.AsMultisig)
	} else if m.IsRecovery {
		err1 = encoder.PushByte(11)
		err2 = encoder.Encode(&m.AsRecovery)
	} else if m.IsProxy {
		err1 = encoder.PushByte(12)
		err2 = encoder.Encode(&m.AsProxy)
	} else if m.IsScheduler {
		err1 = encoder.PushByte(13)
		err2 = encoder.Encode(&m.AsScheduler)
	} else if m.IsIndices {
		err1 = encoder.PushByte(14)
		err2 = encoder.Encode(&m.AsIndices)
	} else if m.IsGraduallyUpdate {
		err1 = encoder.PushByte(15)
		err2 = encoder.Encode(&m.AsGraduallyUpdate)
	} else if m.IsAuthorship {
		err1 = encoder.PushByte(16)
		err2 = encoder.Encode(&m.AsAuthorship)
	} else if m.IsBabe {
		err1 = encoder.PushByte(17)
		err2 = encoder.Encode(&m.AsBabe)
	} else if m.IsGrandpa {
		err1 = encoder.PushByte(18)
		err2 = encoder.Encode(&m.AsGrandpa)
	} else if m.IsStaking {
		err1 = encoder.PushByte(19)
		err2 = encoder.Encode(&m.AsStaking)
	} else if m.IsSession {
		err1 = encoder.PushByte(20)
		err2 = encoder.Encode(&m.AsSession)
	} else if m.IsHistorical {
		err1 = encoder.PushByte(21)
		err2 = encoder.Encode(&m.AsHistorical)
	} else if m.IsGeneralCouncil {
		err1 = encoder.PushByte(22)
		err2 = encoder.Encode(&m.AsGeneralCouncil)
	} else if m.IsGeneralCouncilMembership {
		err1 = encoder.PushByte(23)
		err2 = encoder.Encode(&m.AsGeneralCouncilMembership)
	} else if m.IsHonzonCouncil {
		err1 = encoder.PushByte(24)
		err2 = encoder.Encode(&m.AsHonzonCouncil)
	} else if m.IsHonzonCouncilMembership {
		err1 = encoder.PushByte(25)
		err2 = encoder.Encode(&m.AsHonzonCouncilMembership)
	} else if m.IsHomaCouncil {
		err1 = encoder.PushByte(26)
		err2 = encoder.Encode(&m.AsHomaCouncil)
	} else if m.IsHomaCouncilMembership {
		err1 = encoder.PushByte(27)
		err2 = encoder.Encode(&m.AsHomaCouncilMembership)
	} else if m.IsTechnicalCommittee {
		err1 = encoder.PushByte(28)
		err2 = encoder.Encode(&m.AsTechnicalCommittee)
	} else if m.IsTechnicalCommitteeMembership {
		err1 = encoder.PushByte(29)
		err2 = encoder.Encode(&m.AsTechnicalCommitteeMembership)
	} else if m.IsAuthority {
		err1 = encoder.PushByte(30)
		err2 = encoder.Encode(&m.AsAuthority)
	} else if m.IsElectionsPhragmen {
		err1 = encoder.PushByte(31)
		err2 = encoder.Encode(&m.AsElectionsPhragmen)
	} else if m.IsAcalaOracle {
		err1 = encoder.PushByte(32)
		err2 = encoder.Encode(&m.AsAcalaOracle)
	} else if m.IsBandOracle {
		err1 = encoder.PushByte(33)
		err2 = encoder.Encode(&m.AsBandOracle)
	} else if m.IsOperatorMembershipAcala {
		err1 = encoder.PushByte(34)
		err2 = encoder.Encode(&m.AsOperatorMembershipAcala)
	} else if m.IsOperatorMembershipBand {
		err1 = encoder.PushByte(35)
		err2 = encoder.Encode(&m.AsOperatorMembershipBand)
	} else if m.IsAuction {
		err1 = encoder.PushByte(36)
		err2 = encoder.Encode(&m.AsAuction)
	} else if m.IsRewards {
		err1 = encoder.PushByte(37)
		err2 = encoder.Encode(&m.AsRewards)
	} else if m.IsOrmlNFT {
		err1 = encoder.PushByte(38)
		err2 = encoder.Encode(&m.AsOrmlNFT)
	} else if m.IsPrices {
		err1 = encoder.PushByte(39)
		err2 = encoder.Encode(&m.AsPrices)
	} else if m.IsDex {
		err1 = encoder.PushByte(40)
		err2 = encoder.Encode(&m.AsDex)
	} else if m.IsAuctionManager {
		err1 = encoder.PushByte(41)
		err2 = encoder.Encode(&m.AsAuctionManager)
	} else if m.IsLoans {
		err1 = encoder.PushByte(42)
		err2 = encoder.Encode(&m.AsLoans)
	} else if m.IsHonzon {
		err1 = encoder.PushByte(43)
		err2 = encoder.Encode(&m.AsHonzon)
	} else if m.IsCdpTreasury {
		err1 = encoder.PushByte(44)
		err2 = encoder.Encode(&m.AsCdpTreasury)
	} else if m.IsCdpEngine {
		err1 = encoder.PushByte(45)
		err2 = encoder.Encode(&m.AsCdpEngine)
	} else if m.IsEmergencyShutdown {
		err1 = encoder.PushByte(46)
		err2 = encoder.Encode(&m.AsEmergencyShutdown)
	} else if m.IsHoma {
		err1 = encoder.PushByte(47)
		err2 = encoder.Encode(&m.AsHoma)
	} else if m.IsNomineesElection {
		err1 = encoder.PushByte(48)
		err2 = encoder.Encode(&m.AsNomineesElection)
	} else if m.IsStakingPool {
		err1 = encoder.PushByte(49)
		err2 = encoder.Encode(&m.AsStakingPool)
	} else if m.IsPolkadotBridge {
		err1 = encoder.PushByte(50)
		err2 = encoder.Encode(&m.AsPolkadotBridge)
	} else if m.IsIncentives {
		err1 = encoder.PushByte(51)
		err2 = encoder.Encode(&m.AsIncentives)
	} else if m.IsAirDrop {
		err1 = encoder.PushByte(52)
		err2 = encoder.Encode(&m.AsAirDrop)
	} else if m.IsNFT {
		err1 = encoder.PushByte(53)
		err2 = encoder.Encode(&m.AsNFT)
	} else if m.IsRenVmBridge {
		err1 = encoder.PushByte(54)
		err2 = encoder.Encode(&m.AsRenVmBridge)
	} else if m.IsContracts {
		err1 = encoder.PushByte(55)
		err2 = encoder.Encode(&m.AsContracts)
	} else if m.IsEVM {
		err1 = encoder.PushByte(56)
		err2 = encoder.Encode(&m.AsEVM)
	} else if m.IsSudo {
		err1 = encoder.PushByte(57)
		err2 = encoder.Encode(&m.AsSudo)
	} else if m.IsTransactionPayment {
		err1 = encoder.PushByte(58)
		err2 = encoder.Encode(&m.AsTransactionPayment)
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}
