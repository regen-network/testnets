package src

type Validator struct {
	ValidatorInfo []ValidatorInfo `json:"validatorInfo"`
}

type ValidatorInfo struct {
	ValAddress string `json:"valAddress"`
	Info       Info   `json:"info"`
}

type Info struct {
	UptimeScore   float64 `json:"uptimeScore"`
	Moniker       string  `json:"moniker"`
	OperatorAddr  string  `json:"operatorAddr"`
	Upgrade1Score int64   `json:"upgrade1Score"`
	Upgrade2Score int64   `json:"upgrade2Score"`
	StartBlock    int64   `json:"startBlock"`
	UptimeCount   int64   `json:"uptimeCount"`
}
