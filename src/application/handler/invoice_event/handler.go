package invoiceevent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/snxl/stark_bank_integration/src/config"
	queueclient "github.com/snxl/stark_bank_integration/src/core/client/queue_client"
	_ "github.com/snxl/stark_bank_integration/src/core/dto/payment"
	webhookDTO "github.com/snxl/stark_bank_integration/src/core/dto/webhook"
)

type InvoiceEventHandler struct {
	queue queueclient.QueueClient[asynq.HandlerFunc]
}

func NewInvoicePaidEventHandler() *InvoiceEventHandler {
	return &InvoiceEventHandler{
		queue: queueclient.NewQueueAsynq(*config.GetAsynq().RedisOpt),
	}
}

// @Summary		Transfer webhook
// @Description	Job notification to registry a transfer
// @Tags		webhook
// @Accept		json
// @Produce		json
// @Param		payload	body		webhook.Request true "body payload"
// @Success		200
// @Failure 	404
// @Failure 	500
// @Router		/webhook/starkbank [post]
func (i *InvoiceEventHandler) Run(ctx *gin.Context) {
	var req webhookDTO.Request

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	err = i.queue.ProcessInvoiceEvent(ctx, req.Event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "enqueued"})
}
