package delivery

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	//v              *validator.Validate
	//cfg            Config
	//useCase        usecase.UseCase
	//metrics        metrics.Metrics
	//es             eventrepo.EventStore
	//balance        *balance.Connector // todo не юзать! это только временно
	//swap           *swap.Connector    // todo не юзать! это только временно
	//captchaManager *captchaManager.CaptchaManager
}

func New(
// cfg Config,
// useCase usecase.UseCase,
// metrics metrics.Metrics,
// balanceConn *balance.Connector,
// swapConn *swap.Connector,
// es eventrepo.EventStore,
// captchaManager *captchaManager.CaptchaManager,
) (*Handler, error) {
	//v, err := validation.New()
	//if err != nil {
	//	return nil, err
	//}

	return &Handler{
		//v:              v,
		//es:             es,
		//cfg:            cfg,
		//useCase:        useCase,
		//metrics:        metrics,
		//balance:        balanceConn,
		//swap:           swapConn,
		//captchaManager: captchaManager,
	}, nil
}

// GetForeignCurrencyCardpayment Gets a list of available currencies for cardpayment UZ/KZ.
//
// @Summary Return a list of currencies.
// @Description Currency list contains Name, Description and CurerensyID.
// @Tags Cardpayment
// @Produce		json
// @Success 200 {object} []models.ForeignCurrency "Successful operation"
// @Failure 500 {object} models.FailCauseResult "Failed"
// @Router /cardpayment/currency/get_foreign_currency [get]
func (h *Handler) GetForeignCurrencyCardpayment(c *fiber.Ctx) error {
	//ctx, span := oteltrace.StartFiberTrace(c)
	//defer span.End()

	//foreignCurrency, err := h.useCase.GetForeignCurrency(ctx)
	//if err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(models.FailCauseResult{
	//		FailCause: err.Error(),
	//	})
	//}

	//return c.JSON(foreignCurrency)
	return c.JSON(nil)
}
