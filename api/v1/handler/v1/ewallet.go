package v1

import (
	"github.com/achmad-dev/simple-ewallet/internal/dto"
	"github.com/achmad-dev/simple-ewallet/internal/pkg"
	"github.com/achmad-dev/simple-ewallet/internal/service"
	"github.com/gofiber/fiber/v2"
)

type EWalletHandler interface {
	AddBalance(c *fiber.Ctx) error
	SubtractBalance(c *fiber.Ctx) error
	GetWallet(c *fiber.Ctx) error
}

type eWalletHandler struct {
	ewalletService service.EWalletService
}

// AddBalance implements EWalletHandler.
func (e *eWalletHandler) AddBalance(c *fiber.Ctx) error {
	var request dto.EWalletAddBalanceDto
	if err := c.BodyParser(&request); err != nil {
		return pkg.ErrorResponse(c, fiber.StatusBadRequest, "invalid request")
	}
	if c.Locals("user_id") == nil {
		return pkg.ErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}
	userID := c.Locals("user_id").(string)
	err := e.ewalletService.AddBalance(userID, request.Amount)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return pkg.SuccessResponse(c, nil, "ewallet balance added successfully")
}

// GetWallet implements EWalletHandler.
func (e *eWalletHandler) GetWallet(c *fiber.Ctx) error {
	if c.Locals("user_id") == nil {
		return pkg.ErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}
	userID := c.Locals("user_id").(string)
	ewallet, err := e.ewalletService.GetWallet(userID)
	if err != nil {
		return pkg.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return pkg.SuccessResponse(c, ewallet, "ewallet retrieved successfully")
}

// SubtractBalance implements EWalletHandler.
func (e *eWalletHandler) SubtractBalance(c *fiber.Ctx) error {
	var request dto.EWalletWithdrawBalanceDto
	if err := c.BodyParser(&request); err != nil {
		return pkg.ErrorResponse(c, fiber.StatusBadRequest, "invalid request")
	}
	if c.Locals("user_id") == nil {
		return pkg.ErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}
	userID := c.Locals("user_id").(string)
	err := e.ewalletService.SubtractBalance(userID, request.Amount)
	if err != nil {
		return pkg.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return pkg.SuccessResponse(c, nil, "ewallet balance withdraw successfully")
}

func NewEWalletHandler(ewalletService service.EWalletService) EWalletHandler {
	return &eWalletHandler{
		ewalletService: ewalletService,
	}
}
