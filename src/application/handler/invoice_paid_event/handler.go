package invoicepaidevent

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snxl/stark_bank_integration/src/core/client/payment"
	_ "github.com/snxl/stark_bank_integration/src/core/dto/payment"
	invoicepaidevent "github.com/snxl/stark_bank_integration/src/core/usecase/invoice_paid_event"
)

type InvoicePaidEventHandler struct{}

func NewInvoicePaidEventHandler() *InvoicePaidEventHandler {
	return &InvoicePaidEventHandler{}
}

// @Summary		Transfer webhook
// @Description	Job notification to registry a transfer
// @Tags		event
// @Accept		json
// @Produce		json
// @Success		200		{object}	payment.SendTransferDTO
// @Router		/webhook [post]
func (i *InvoicePaidEventHandler) Run(ctx *gin.Context) {
	type Invoice struct {
		Amount int    `json:",omitempty"`
		Status string `json:",omitempty"`
	}

	type Log struct {
		Invoice Invoice `json:"invoice"`
	}

	type request struct {
		Event struct {
			Subscription string `json:"subscription"`
			Log          Log    `json:"log"`
		} `json:"event"`
		// Event struct {
		// 	Subscription string `json:"subscription"`
		// 	Log          struct {
		// 		Invoice struct {
		// 			Amount int `json:"amount"`
		// 		} `json:"invoice"`
		// 	} `json:"log"`
		// } `json:"event"`
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	fmt.Println(req)
	usecase := invoicepaidevent.NewInvoicePaidEventUsecase(
		payment.NewStarkbankSDKClient(),
	)
	err := usecase.Run(invoicepaidevent.Input{
		Subscription: req.Event.Subscription,
		Status:       req.Event.Log.Invoice.Status,
		Amount:       req.Event.Log.Invoice.Amount,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, nil)
}
