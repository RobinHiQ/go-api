package api

import (
	"net/http"

	functions "github.com/RobinHiQ/go-api/functions"
	"github.com/gin-gonic/gin"
)

//	@Summary		Generate job description
//	@Description	Generates a job description for the given job title and saves it to Redis
//	@Param			jobTitle	query	string	true	"The title of the job"
//	@Produce		json
//	@Success		200	{string}	string	"Successful operation"
//	@Router			/job-description [get]
func GetJobDescription(c *gin.Context) {
	jobTitle := c.Query("jobTitle")
	if jobTitle == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing jobTitle parameter"})
		return
	}

	description, err := functions.GenerateJobDescription(jobTitle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, description)
}
