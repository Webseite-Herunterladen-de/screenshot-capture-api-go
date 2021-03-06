// Code generated by go-swagger; DO NOT EDIT.

package screenshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/Webseite-Herunterladen-de/screenshot-capture-api-go/models"
)

// CaptureScreenshotAuthenticatedReader is a Reader for the CaptureScreenshotAuthenticated structure.
type CaptureScreenshotAuthenticatedReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *CaptureScreenshotAuthenticatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCaptureScreenshotAuthenticatedOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 302:
		result := NewCaptureScreenshotAuthenticatedFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCaptureScreenshotAuthenticatedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCaptureScreenshotAuthenticatedOK creates a CaptureScreenshotAuthenticatedOK with default headers values
func NewCaptureScreenshotAuthenticatedOK(writer io.Writer) *CaptureScreenshotAuthenticatedOK {
	return &CaptureScreenshotAuthenticatedOK{

		Payload: writer,
	}
}

/* CaptureScreenshotAuthenticatedOK describes a response with status code 200, with default header values.

the requested screenshot as binary
*/
type CaptureScreenshotAuthenticatedOK struct {
	ContentType            string
	Location               string
	XREMAININGQUOTA        int64
	XREMAININGQUOTAPREPAID int64

	Payload io.Writer
}

func (o *CaptureScreenshotAuthenticatedOK) Error() string {
	return fmt.Sprintf("[GET /capture/{token}/{hash}][%d] captureScreenshotAuthenticatedOK  %+v", 200, o.Payload)
}
func (o *CaptureScreenshotAuthenticatedOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *CaptureScreenshotAuthenticatedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Type
	hdrContentType := response.GetHeader("Content-Type")

	if hdrContentType != "" {
		o.ContentType = hdrContentType
	}

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	// hydrates response header X-REMAINING-QUOTA
	hdrXREMAININGQUOTA := response.GetHeader("X-REMAINING-QUOTA")

	if hdrXREMAININGQUOTA != "" {
		valxREMAININGQUOTA, err := swag.ConvertInt64(hdrXREMAININGQUOTA)
		if err != nil {
			return errors.InvalidType("X-REMAINING-QUOTA", "header", "int64", hdrXREMAININGQUOTA)
		}
		o.XREMAININGQUOTA = valxREMAININGQUOTA
	}

	// hydrates response header X-REMAINING-QUOTA-PREPAID
	hdrXREMAININGQUOTAPREPAID := response.GetHeader("X-REMAINING-QUOTA-PREPAID")

	if hdrXREMAININGQUOTAPREPAID != "" {
		valxREMAININGQUOTAPREPAId, err := swag.ConvertInt64(hdrXREMAININGQUOTAPREPAID)
		if err != nil {
			return errors.InvalidType("X-REMAINING-QUOTA-PREPAID", "header", "int64", hdrXREMAININGQUOTAPREPAID)
		}
		o.XREMAININGQUOTAPREPAID = valxREMAININGQUOTAPREPAId
	}

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCaptureScreenshotAuthenticatedFound creates a CaptureScreenshotAuthenticatedFound with default headers values
func NewCaptureScreenshotAuthenticatedFound() *CaptureScreenshotAuthenticatedFound {
	return &CaptureScreenshotAuthenticatedFound{}
}

/* CaptureScreenshotAuthenticatedFound describes a response with status code 302, with default header values.

the requested screenshot as redirect
*/
type CaptureScreenshotAuthenticatedFound struct {
	Location               string
	XREMAININGQUOTA        int64
	XREMAININGQUOTAPREPAID int64
}

func (o *CaptureScreenshotAuthenticatedFound) Error() string {
	return fmt.Sprintf("[GET /capture/{token}/{hash}][%d] captureScreenshotAuthenticatedFound ", 302)
}

func (o *CaptureScreenshotAuthenticatedFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	// hydrates response header X-REMAINING-QUOTA
	hdrXREMAININGQUOTA := response.GetHeader("X-REMAINING-QUOTA")

	if hdrXREMAININGQUOTA != "" {
		valxREMAININGQUOTA, err := swag.ConvertInt64(hdrXREMAININGQUOTA)
		if err != nil {
			return errors.InvalidType("X-REMAINING-QUOTA", "header", "int64", hdrXREMAININGQUOTA)
		}
		o.XREMAININGQUOTA = valxREMAININGQUOTA
	}

	// hydrates response header X-REMAINING-QUOTA-PREPAID
	hdrXREMAININGQUOTAPREPAID := response.GetHeader("X-REMAINING-QUOTA-PREPAID")

	if hdrXREMAININGQUOTAPREPAID != "" {
		valxREMAININGQUOTAPREPAId, err := swag.ConvertInt64(hdrXREMAININGQUOTAPREPAID)
		if err != nil {
			return errors.InvalidType("X-REMAINING-QUOTA-PREPAID", "header", "int64", hdrXREMAININGQUOTAPREPAID)
		}
		o.XREMAININGQUOTAPREPAID = valxREMAININGQUOTAPREPAId
	}

	return nil
}

// NewCaptureScreenshotAuthenticatedDefault creates a CaptureScreenshotAuthenticatedDefault with default headers values
func NewCaptureScreenshotAuthenticatedDefault(code int) *CaptureScreenshotAuthenticatedDefault {
	return &CaptureScreenshotAuthenticatedDefault{
		_statusCode: code,
	}
}

/* CaptureScreenshotAuthenticatedDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CaptureScreenshotAuthenticatedDefault struct {
	_statusCode int
	ContentType string

	Payload *models.ErrorModel
}

// Code gets the status code for the capture screenshot authenticated default response
func (o *CaptureScreenshotAuthenticatedDefault) Code() int {
	return o._statusCode
}

func (o *CaptureScreenshotAuthenticatedDefault) Error() string {
	return fmt.Sprintf("[GET /capture/{token}/{hash}][%d] captureScreenshotAuthenticated default  %+v", o._statusCode, o.Payload)
}
func (o *CaptureScreenshotAuthenticatedDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CaptureScreenshotAuthenticatedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Type
	hdrContentType := response.GetHeader("Content-Type")

	if hdrContentType != "" {
		o.ContentType = hdrContentType
	}

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
