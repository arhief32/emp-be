package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/arhief32/emp-be/v1/entities"
	"github.com/arhief32/emp-be/v1/services"
)

type MerchantSubmissionController struct {
	svc services.MerchantSubmissionService
}

func NewMerchantSubmissionController(s services.MerchantSubmissionService) MerchantSubmissionController {
	return MerchantSubmissionController{svc: s}
}

// Create draft (Maker)
func (ctr *MerchantSubmissionController) Create(c *gin.Context) {
	var req entities.MerchantSubmissionCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// read maker id from context
	makerIDRaw, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	makerID := uint(makerIDRaw.(int))

	// convert to map for service helper
	m := map[string]interface{}{
		"merchant_name":    req.MerchantName,
		"owner_name":       req.OwnerName,
		"phone":            req.Phone,
		"email":            req.Email,
		"address":          req.Address,
		"category":         req.Category,
		"nib":              req.NIB,
		"npwp":             req.NPWP,
		"year_established": req.YearEstablished,
		"employees":        req.Employees,
		"documents":        req.Documents,
	}
	sub, err := ctr.svc.Create(m, makerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": sub})
}

// Update draft
func (ctr *MerchantSubmissionController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id64, _ := strconv.ParseUint(idStr, 10, 64)
	id := uint(id64)

	var req entities.MerchantSubmissionUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	makerIDRaw, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	makerID := uint(makerIDRaw.(int))

	// convert to map
	payload := map[string]interface{}{}
	if req.MerchantName != nil {
		payload["merchant_name"] = *req.MerchantName
	}
	if req.OwnerName != nil {
		payload["owner_name"] = *req.OwnerName
	}
	if req.Phone != nil {
		payload["phone"] = *req.Phone
	}
	if req.Email != nil {
		payload["email"] = *req.Email
	}
	if req.Address != nil {
		payload["address"] = *req.Address
	}
	if req.Category != nil {
		payload["category"] = *req.Category
	}
	if req.NIB != nil {
		payload["nib"] = *req.NIB
	}
	if req.NPWP != nil {
		payload["npwp"] = *req.NPWP
	}
	if req.YearEstablished != nil {
		payload["year_established"] = *req.YearEstablished
	}
	if req.Employees != nil {
		payload["employees"] = *req.Employees
	}
	if req.Documents != nil {
		payload["documents"] = req.Documents
	}

	sub, err := ctr.svc.Update(id, payload, makerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sub})
}

// Submit (Maker -> Checker)
func (ctr *MerchantSubmissionController) Submit(c *gin.Context) {
	idStr := c.Param("id")
	id64, _ := strconv.ParseUint(idStr, 10, 64)
	id := uint(id64)

	makerIDRaw, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	makerID := uint(makerIDRaw.(int))

	if err := ctr.svc.Submit(id, makerID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "submitted"})
}

// List own submissions (Maker)
func (ctr *MerchantSubmissionController) ListMine(c *gin.Context) {
	makerIDRaw, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	makerID := uint(makerIDRaw.(int))
	list, err := ctr.svc.GetByMaker(makerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

// Checker: list pending
func (ctr *MerchantSubmissionController) ListPendingForChecker(c *gin.Context) {
	list, err := ctr.svc.ListPendingForChecker()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

// Checker approve
func (ctr *MerchantSubmissionController) CheckerApprove(c *gin.Context) {
	idStr := c.Param("id")
	id64, _ := strconv.ParseUint(idStr, 10, 64)
	id := uint(id64)

	checkerIDRaw, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	checkerID := uint(checkerIDRaw.(int))

	var body struct {
		Notes string `json:"notes"`
	}
	_ = c.ShouldBindJSON(&body)

	if err := ctr.svc.CheckerApprove(id, checkerID, body.Notes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "checked"})
}

// Checker reject
func (ctr *MerchantSubmissionController) CheckerReject(c *gin.Context) {
	idStr := c.Param("id")
	id64, _ := strconv.ParseUint(idStr, 10, 64)
	id := uint(id64)
	checkerIDRaw, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	checkerID := uint(checkerIDRaw.(int))

	var body struct {
		Notes string `json:"notes" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctr.svc.CheckerReject(id, checkerID, body.Notes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "rejected"})
}

// Signer approve
func (ctr *MerchantSubmissionController) SignerApprove(c *gin.Context) {
	idStr := c.Param("id")
	id64, _ := strconv.ParseUint(idStr, 10, 64)
	id := uint(id64)

	signerIDRaw, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	signerID := uint(signerIDRaw.(int))

	if err := ctr.svc.SignerApprove(id, signerID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "approved and merchant created"})
}

// Signer reject
func (ctr *MerchantSubmissionController) SignerReject(c *gin.Context) {
	idStr := c.Param("id")
	id64, _ := strconv.ParseUint(idStr, 10, 64)
	id := uint(id64)

	signerIDRaw, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	signerID := uint(signerIDRaw.(int))

	var body struct {
		Notes string `json:"notes" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctr.svc.SignerReject(id, signerID, body.Notes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "rejected"})
}
