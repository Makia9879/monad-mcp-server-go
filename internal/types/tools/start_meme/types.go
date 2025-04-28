package start_meme

type StartMemeRequest struct {
	MemePumpURL string `json:"meme_pump_url" description:"meme发射平台的地址，如果没有填写则默认值为 https://testnet.nad.fun/" required:"false"`
}
