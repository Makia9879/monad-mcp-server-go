package quick_trade

type QuickTradeRequest struct {
	TargetAbstract string  `json:"target_abstract" description:"target abstract" required:"true"`
	SlipPoint      float32 `json:"slip_point" description:"slip point, maximum is 90.0, minimum is 0.1" default:"30.0" required:"false"`
	Pay            float32 `json:"pay" required:"false" description:"how much MON will be pay for, its value will be from pay_percent if not set"`
	PayPercent     float32 `json:"pay_percent" required:"false" description:"Pay the balance ratio of MON in wallet, maximum is 100.0, minimum is 0.0" default:"50.0"`
}
