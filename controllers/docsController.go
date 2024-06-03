package controllers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	storage_go "github.com/supabase-community/storage-go"
)

func handleError(c *gin.Context, err error, status int, message string) {
	if err != nil {
		c.JSON(status, gin.H{"message": message, "error": err.Error()})
		return
	}
}

func GetFormDocs(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

func UploadFiles(c *gin.Context) {
	// Connect to the storage bucket
	API_SECRET_KEY := os.Getenv("API_SECRET_KEY")
	storageClient := storage_go.NewClient("https://lwqkmzjmtcjagxkwvosj.supabase.co/storage/v1", API_SECRET_KEY, nil)

	var errs []error
	var files map[string]multipart.File
	formId := c.Param("id")
	loanNumber := c.Query("loan_number")

	// Url For Redirect
	targetURL := "http://localhost:5000?success=true"

	fmt.Printf("loanNumber: %v\n", loanNumber)
	fmt.Printf("formId: %v\n", formId)

	switch formId {
	case "returning":
		// Extract (pdf) files from request
		nrc, _, err1 := c.Request.FormFile("nrc")
		confirmation_slip, _, err2 := c.Request.FormFile("confirmation_slip")
		payment_history, _, err3 := c.Request.FormFile("payment_history")
		bank_statement, _, err4 := c.Request.FormFile("bank_statement")
		transcript, _, err5 := c.Request.FormFile("transcript")

		errs = append(errs, err1, err2, err3, err4, err5)

		files = map[string]multipart.File{
			"nrc":               nrc,
			"confirmation_slip": confirmation_slip,
			"payment_history":   payment_history,
			"bank_statement":    bank_statement,
			"transcript":        transcript,
		}

	case "first":
		// Extract (pdf) files from request
		nrc, _, err1 := c.Request.FormFile("nrc")
		guardian_nrc, _, err2 := c.Request.FormFile("guardian_nrc")
		confirmation_slip, _, err3 := c.Request.FormFile("confirmation_slip")
		payment_history, _, err4 := c.Request.FormFile("payment_history")
		bank_statement, _, err5 := c.Request.FormFile("bank_statement")
		tpin, _, err6 := c.Request.FormFile("tpin")
		transcript, _, err7 := c.Request.FormFile("transcript")

		errs = append(errs, err1, err2, err3, err4, err5, err6, err7)

		files = map[string]multipart.File{
			"nrc":               nrc,
			"guardian_nrc":      guardian_nrc,
			"confirmation_slip": confirmation_slip,
			"payment_history":   payment_history,
			"bank_statement":    bank_statement,
			"tpin":              tpin,
			"transcript":        transcript,
		}

		// Add loan number to the redirect url
		targetURL += "&loan_number=" + loanNumber
	}

	// Handle Retrieval errors for every file
	for _, err := range errs {
		handleError(c, err, http.StatusBadRequest, "Failed to Load PDF")
	}

	// Upload every file
	for fileName, file := range files {
		// Read the file content into a byte slice
		fileContent, err := io.ReadAll(file)
		handleError(c, err, http.StatusInternalServerError, "Failed to Read PDF file")

		// Convert that byte slice to type of `Reader`.
		fileBody := bytes.NewReader(fileContent)

		// Now we can use `fileBody` for uploading to Supabase.
		// ...
		storagePath := loanNumber + "/" + fileName + ".pdf"

		if formId == "first" {
			// Upload Files
			result, err := storageClient.UploadFile("docs", storagePath, fileBody)
			handleError(c, err, http.StatusInternalServerError, "Failed to Upload pdf file: "+fileName)
			fmt.Printf("uploaded: %v\n", result.Key)
		} else {
			// Update Files
			result, err := storageClient.UpdateFile("docs", storagePath, fileBody)
			handleError(c, err, http.StatusInternalServerError, "Failed to Update pdf file: "+fileName)
			fmt.Printf("uploaded: %v\n", result.Key)
		}
	}
	// Redirect the user to the target URL
	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, targetURL)

}

func Redirect(c *gin.Context) {
	loanNumber := c.Query("loan_number")
	id := c.Param("id")
	targetURL := "http://localhost:5000/?success=true"

	// Add query parameters
	if id == "first" {
		targetURL += "&loan_number=" + loanNumber
	}
	c.Redirect(http.StatusPermanentRedirect, targetURL)
}
