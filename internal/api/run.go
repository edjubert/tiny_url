package api

func (api *Api) Run() error {
	api.Logger.Info().Msg("API running on port " + api.Config.Server.Port)
	return api.Server.Run(":" + api.Config.Server.Port)
}
