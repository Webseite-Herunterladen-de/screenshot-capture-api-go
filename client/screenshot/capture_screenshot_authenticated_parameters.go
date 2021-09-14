// Code generated by go-swagger; DO NOT EDIT.

package screenshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewCaptureScreenshotAuthenticatedParams creates a new CaptureScreenshotAuthenticatedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCaptureScreenshotAuthenticatedParams() *CaptureScreenshotAuthenticatedParams {
	return &CaptureScreenshotAuthenticatedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCaptureScreenshotAuthenticatedParamsWithTimeout creates a new CaptureScreenshotAuthenticatedParams object
// with the ability to set a timeout on a request.
func NewCaptureScreenshotAuthenticatedParamsWithTimeout(timeout time.Duration) *CaptureScreenshotAuthenticatedParams {
	return &CaptureScreenshotAuthenticatedParams{
		timeout: timeout,
	}
}

// NewCaptureScreenshotAuthenticatedParamsWithContext creates a new CaptureScreenshotAuthenticatedParams object
// with the ability to set a context for a request.
func NewCaptureScreenshotAuthenticatedParamsWithContext(ctx context.Context) *CaptureScreenshotAuthenticatedParams {
	return &CaptureScreenshotAuthenticatedParams{
		Context: ctx,
	}
}

// NewCaptureScreenshotAuthenticatedParamsWithHTTPClient creates a new CaptureScreenshotAuthenticatedParams object
// with the ability to set a custom HTTPClient for a request.
func NewCaptureScreenshotAuthenticatedParamsWithHTTPClient(client *http.Client) *CaptureScreenshotAuthenticatedParams {
	return &CaptureScreenshotAuthenticatedParams{
		HTTPClient: client,
	}
}

/* CaptureScreenshotAuthenticatedParams contains all the parameters to send to the API endpoint
   for the capture screenshot authenticated operation.

   Typically these are written to a http.Request.
*/
type CaptureScreenshotAuthenticatedParams struct {

	/* Accuracy.

	   The accuracy in meters used in the emulation of the geo-location.

	   Format: float64
	   Default: 2
	*/
	Accuracy *float64

	/* Adblock.

	   Prevent ads from being displayed. Block requests from popular ad networks and hide frequent ads.
	*/
	Adblock *bool

	/* Cookies.

	   A semicolon-separated list of cookies to be used when capturing the screenshot. Each cookies should be passed as a key-value pair and multiple pairs should be separated by a semicolon.
	*/
	Cookies *string

	/* CSS.

	   Inject your custom CSS.
	*/
	CSS *string

	/* Delay.

	   The delay in milliseconds to wait after the page loads before taking the screenshot. This is in milliseconds. One second is 1000 milliseconds. From 0 milliseconds to a maximum of 10,000 milliseconds.

	   Format: int64
	*/
	Delay *int64

	/* Device.

	   The device used in the emulation.
	*/
	Device *string

	/* Element.

	   Takes a screenshot of the first element matched by the specified CSS selector. This is ignored if full is true. (This option cannot be used with the PDF export format.)
	*/
	Element *string

	/* FileType.

	   The image file format of the captured screenshot. Either png, jpeg, webp or PDF with 72 dpi.

	   Default: "png"
	*/
	FileType *string

	/* Full.

	   Set this parameter to true if you want to screenshot the whole web page in full size.
	*/
	Full *bool

	/* Hash.

	   The hash value is for authenticated requests. If you want to publish this URL, you should use the authenticated requests.
	*/
	Hash string

	/* Headers.

	   A semicolon-separated list of header parameters to be used when capturing the screenshot. Each header should be passed as a key-value pair and multiple pairs should be separated by a semicolon.
	*/
	Headers *string

	/* Height.

	   The height, in pixels, of the browser viewport to use. Ignored if you set full to true.

	   Format: int64
	   Default: 1080
	*/
	Height *int64

	/* HideCookieBanners.

	   Prevent cookie banners and pop-ups from being displayed. The best possible result is tried.
	*/
	HideCookieBanners *bool

	/* Invalidate.

	   Force the API to invalidate the cache and capture a new screenshot. This call costs you additional money, because a call of a cache hit is not charged.
	*/
	Invalidate *bool

	/* Js.

	   Inject your custom Javascript.
	*/
	Js *string

	/* Language.

	   Sets the Accept-Language header on requests to the target URL so that you can take screenshots from a website with a specific language.
	*/
	Language *string

	/* Latitude.

	   The latitude used in the emulation of the geo-location.

	   Format: float64
	*/
	Latitude *float64

	/* LazyloadScroll.

	   Set this parameter to true to scroll down through the entire page before taking a screenshot. This is useful for triggering animations or lazy load elements in full screen.
	*/
	LazyloadScroll *bool

	/* Longitude.

	   The longitude used in the emulation of the geo-location.

	   Format: float64
	*/
	Longitude *float64

	/* Proxy.

	   Use an address of a proxy server through which the screenshot should be taken. The proxy address should be formatted as http://username:password@proxyserver.com:31280
	*/
	Proxy *string

	/* Quality.

	   The quality of the image between 0 and 100. This works only for the jpeg format, for PNG images the parameter is applied only during compression.

	   Format: int64
	   Default: 90
	*/
	Quality *int64

	/* RandomUserAgent.

	   Sets a random user agent header to emulate a different devices when taking screenshots.
	*/
	RandomUserAgent *bool

	/* Redirect.

	   If you set Redirect, the response will be a 302 redirect to the screenshot file in our CDN.
	*/
	Redirect *bool

	/* Scale.

	   The scale factor of the device to use when taking the screenshot. For example, a scale factor of 2 produces a high-resolution screenshot suitable for viewing on Retina devices. The larger the scale factor, the larger the screenshot produced.

	   Format: float64
	   Default: 1
	*/
	Scale *float64

	/* Timezone.

	   The IANA time zone identifier used for this capture.

	   Default: "Europe/Berlin"
	*/
	Timezone *string

	/* Token.

	   A valid token is needed to make paid API calls. Tokens can be managed from your account.
	*/
	Token string

	/* TTL.

	   Number of seconds the capture file is cached by our CDN. An API request that is loaded through the cache does not count as a paid request. You can set a number of seconds from 0 seconds up to 2592000 seconds. This is a maximum of 30 days.

	   Format: int64
	*/
	TTL *int64

	/* URL.

	   The URL of the website you want to capture. Please include the protocol (http:// or https://).
	*/
	URL string

	/* UserAgent.

	   Sets the user agent header to emulate a specific device when taking screenshots.
	*/
	UserAgent *string

	/* Wait.

	   Wait until the specified CSS selector matches an element present in the page before taking a screenshot. The process is canceled after 60 seconds.
	*/
	Wait *string

	/* Width.

	   The width, in pixels, of the browser viewport to use.

	   Format: int64
	   Default: 1920
	*/
	Width *int64

	/* X.

	   The starting point of a section screenshot on the X axis.

	   Format: int64
	*/
	X *int64

	/* Y.

	   The starting point of a section screenshot on the Y axis.

	   Format: int64
	*/
	Y *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the capture screenshot authenticated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CaptureScreenshotAuthenticatedParams) WithDefaults() *CaptureScreenshotAuthenticatedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the capture screenshot authenticated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CaptureScreenshotAuthenticatedParams) SetDefaults() {
	var (
		accuracyDefault = float64(2)

		adblockDefault = bool(false)

		fileTypeDefault = string("png")

		heightDefault = int64(1080)

		hideCookieBannersDefault = bool(false)

		latitudeDefault = float64(0)

		lazyloadScrollDefault = bool(false)

		longitudeDefault = float64(0)

		qualityDefault = int64(90)

		randomUserAgentDefault = bool(false)

		redirectDefault = bool(false)

		scaleDefault = float64(1)

		timezoneDefault = string("Europe/Berlin")

		widthDefault = int64(1920)

		xDefault = int64(0)

		yDefault = int64(0)
	)

	val := CaptureScreenshotAuthenticatedParams{
		Accuracy:          &accuracyDefault,
		Adblock:           &adblockDefault,
		FileType:          &fileTypeDefault,
		Height:            &heightDefault,
		HideCookieBanners: &hideCookieBannersDefault,
		Latitude:          &latitudeDefault,
		LazyloadScroll:    &lazyloadScrollDefault,
		Longitude:         &longitudeDefault,
		Quality:           &qualityDefault,
		RandomUserAgent:   &randomUserAgentDefault,
		Redirect:          &redirectDefault,
		Scale:             &scaleDefault,
		Timezone:          &timezoneDefault,
		Width:             &widthDefault,
		X:                 &xDefault,
		Y:                 &yDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithTimeout(timeout time.Duration) *CaptureScreenshotAuthenticatedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithContext(ctx context.Context) *CaptureScreenshotAuthenticatedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithHTTPClient(client *http.Client) *CaptureScreenshotAuthenticatedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccuracy adds the accuracy to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithAccuracy(accuracy *float64) *CaptureScreenshotAuthenticatedParams {
	o.SetAccuracy(accuracy)
	return o
}

// SetAccuracy adds the accuracy to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetAccuracy(accuracy *float64) {
	o.Accuracy = accuracy
}

// WithAdblock adds the adblock to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithAdblock(adblock *bool) *CaptureScreenshotAuthenticatedParams {
	o.SetAdblock(adblock)
	return o
}

// SetAdblock adds the adblock to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetAdblock(adblock *bool) {
	o.Adblock = adblock
}

// WithCookies adds the cookies to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithCookies(cookies *string) *CaptureScreenshotAuthenticatedParams {
	o.SetCookies(cookies)
	return o
}

// SetCookies adds the cookies to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetCookies(cookies *string) {
	o.Cookies = cookies
}

// WithCSS adds the css to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithCSS(css *string) *CaptureScreenshotAuthenticatedParams {
	o.SetCSS(css)
	return o
}

// SetCSS adds the css to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetCSS(css *string) {
	o.CSS = css
}

// WithDelay adds the delay to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithDelay(delay *int64) *CaptureScreenshotAuthenticatedParams {
	o.SetDelay(delay)
	return o
}

// SetDelay adds the delay to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetDelay(delay *int64) {
	o.Delay = delay
}

// WithDevice adds the device to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithDevice(device *string) *CaptureScreenshotAuthenticatedParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetDevice(device *string) {
	o.Device = device
}

// WithElement adds the element to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithElement(element *string) *CaptureScreenshotAuthenticatedParams {
	o.SetElement(element)
	return o
}

// SetElement adds the element to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetElement(element *string) {
	o.Element = element
}

// WithFileType adds the fileType to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithFileType(fileType *string) *CaptureScreenshotAuthenticatedParams {
	o.SetFileType(fileType)
	return o
}

// SetFileType adds the fileType to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetFileType(fileType *string) {
	o.FileType = fileType
}

// WithFull adds the full to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithFull(full *bool) *CaptureScreenshotAuthenticatedParams {
	o.SetFull(full)
	return o
}

// SetFull adds the full to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetFull(full *bool) {
	o.Full = full
}

// WithHash adds the hash to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithHash(hash string) *CaptureScreenshotAuthenticatedParams {
	o.SetHash(hash)
	return o
}

// SetHash adds the hash to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetHash(hash string) {
	o.Hash = hash
}

// WithHeaders adds the headers to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithHeaders(headers *string) *CaptureScreenshotAuthenticatedParams {
	o.SetHeaders(headers)
	return o
}

// SetHeaders adds the headers to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetHeaders(headers *string) {
	o.Headers = headers
}

// WithHeight adds the height to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithHeight(height *int64) *CaptureScreenshotAuthenticatedParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetHeight(height *int64) {
	o.Height = height
}

// WithHideCookieBanners adds the hideCookieBanners to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithHideCookieBanners(hideCookieBanners *bool) *CaptureScreenshotAuthenticatedParams {
	o.SetHideCookieBanners(hideCookieBanners)
	return o
}

// SetHideCookieBanners adds the hideCookieBanners to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetHideCookieBanners(hideCookieBanners *bool) {
	o.HideCookieBanners = hideCookieBanners
}

// WithInvalidate adds the invalidate to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithInvalidate(invalidate *bool) *CaptureScreenshotAuthenticatedParams {
	o.SetInvalidate(invalidate)
	return o
}

// SetInvalidate adds the invalidate to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetInvalidate(invalidate *bool) {
	o.Invalidate = invalidate
}

// WithJs adds the js to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithJs(js *string) *CaptureScreenshotAuthenticatedParams {
	o.SetJs(js)
	return o
}

// SetJs adds the js to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetJs(js *string) {
	o.Js = js
}

// WithLanguage adds the language to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithLanguage(language *string) *CaptureScreenshotAuthenticatedParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetLanguage(language *string) {
	o.Language = language
}

// WithLatitude adds the latitude to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithLatitude(latitude *float64) *CaptureScreenshotAuthenticatedParams {
	o.SetLatitude(latitude)
	return o
}

// SetLatitude adds the latitude to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetLatitude(latitude *float64) {
	o.Latitude = latitude
}

// WithLazyloadScroll adds the lazyloadScroll to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithLazyloadScroll(lazyloadScroll *bool) *CaptureScreenshotAuthenticatedParams {
	o.SetLazyloadScroll(lazyloadScroll)
	return o
}

// SetLazyloadScroll adds the lazyloadScroll to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetLazyloadScroll(lazyloadScroll *bool) {
	o.LazyloadScroll = lazyloadScroll
}

// WithLongitude adds the longitude to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithLongitude(longitude *float64) *CaptureScreenshotAuthenticatedParams {
	o.SetLongitude(longitude)
	return o
}

// SetLongitude adds the longitude to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetLongitude(longitude *float64) {
	o.Longitude = longitude
}

// WithProxy adds the proxy to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithProxy(proxy *string) *CaptureScreenshotAuthenticatedParams {
	o.SetProxy(proxy)
	return o
}

// SetProxy adds the proxy to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetProxy(proxy *string) {
	o.Proxy = proxy
}

// WithQuality adds the quality to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithQuality(quality *int64) *CaptureScreenshotAuthenticatedParams {
	o.SetQuality(quality)
	return o
}

// SetQuality adds the quality to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetQuality(quality *int64) {
	o.Quality = quality
}

// WithRandomUserAgent adds the randomUserAgent to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithRandomUserAgent(randomUserAgent *bool) *CaptureScreenshotAuthenticatedParams {
	o.SetRandomUserAgent(randomUserAgent)
	return o
}

// SetRandomUserAgent adds the randomUserAgent to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetRandomUserAgent(randomUserAgent *bool) {
	o.RandomUserAgent = randomUserAgent
}

// WithRedirect adds the redirect to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithRedirect(redirect *bool) *CaptureScreenshotAuthenticatedParams {
	o.SetRedirect(redirect)
	return o
}

// SetRedirect adds the redirect to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetRedirect(redirect *bool) {
	o.Redirect = redirect
}

// WithScale adds the scale to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithScale(scale *float64) *CaptureScreenshotAuthenticatedParams {
	o.SetScale(scale)
	return o
}

// SetScale adds the scale to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetScale(scale *float64) {
	o.Scale = scale
}

// WithTimezone adds the timezone to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithTimezone(timezone *string) *CaptureScreenshotAuthenticatedParams {
	o.SetTimezone(timezone)
	return o
}

// SetTimezone adds the timezone to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetTimezone(timezone *string) {
	o.Timezone = timezone
}

// WithToken adds the token to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithToken(token string) *CaptureScreenshotAuthenticatedParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetToken(token string) {
	o.Token = token
}

// WithTTL adds the ttl to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithTTL(ttl *int64) *CaptureScreenshotAuthenticatedParams {
	o.SetTTL(ttl)
	return o
}

// SetTTL adds the ttl to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetTTL(ttl *int64) {
	o.TTL = ttl
}

// WithURL adds the url to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithURL(url string) *CaptureScreenshotAuthenticatedParams {
	o.SetURL(url)
	return o
}

// SetURL adds the url to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetURL(url string) {
	o.URL = url
}

// WithUserAgent adds the userAgent to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithUserAgent(userAgent *string) *CaptureScreenshotAuthenticatedParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithWait adds the wait to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithWait(wait *string) *CaptureScreenshotAuthenticatedParams {
	o.SetWait(wait)
	return o
}

// SetWait adds the wait to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetWait(wait *string) {
	o.Wait = wait
}

// WithWidth adds the width to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithWidth(width *int64) *CaptureScreenshotAuthenticatedParams {
	o.SetWidth(width)
	return o
}

// SetWidth adds the width to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetWidth(width *int64) {
	o.Width = width
}

// WithX adds the x to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithX(x *int64) *CaptureScreenshotAuthenticatedParams {
	o.SetX(x)
	return o
}

// SetX adds the x to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetX(x *int64) {
	o.X = x
}

// WithY adds the y to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) WithY(y *int64) *CaptureScreenshotAuthenticatedParams {
	o.SetY(y)
	return o
}

// SetY adds the y to the capture screenshot authenticated params
func (o *CaptureScreenshotAuthenticatedParams) SetY(y *int64) {
	o.Y = y
}

// WriteToRequest writes these params to a swagger request
func (o *CaptureScreenshotAuthenticatedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Accuracy != nil {

		// query param accuracy
		var qrAccuracy float64

		if o.Accuracy != nil {
			qrAccuracy = *o.Accuracy
		}
		qAccuracy := swag.FormatFloat64(qrAccuracy)
		if qAccuracy != "" {

			if err := r.SetQueryParam("accuracy", qAccuracy); err != nil {
				return err
			}
		}
	}

	if o.Adblock != nil {

		// query param adblock
		var qrAdblock bool

		if o.Adblock != nil {
			qrAdblock = *o.Adblock
		}
		qAdblock := swag.FormatBool(qrAdblock)
		if qAdblock != "" {

			if err := r.SetQueryParam("adblock", qAdblock); err != nil {
				return err
			}
		}
	}

	if o.Cookies != nil {

		// query param cookies
		var qrCookies string

		if o.Cookies != nil {
			qrCookies = *o.Cookies
		}
		qCookies := qrCookies
		if qCookies != "" {

			if err := r.SetQueryParam("cookies", qCookies); err != nil {
				return err
			}
		}
	}

	if o.CSS != nil {

		// query param css
		var qrCSS string

		if o.CSS != nil {
			qrCSS = *o.CSS
		}
		qCSS := qrCSS
		if qCSS != "" {

			if err := r.SetQueryParam("css", qCSS); err != nil {
				return err
			}
		}
	}

	if o.Delay != nil {

		// query param delay
		var qrDelay int64

		if o.Delay != nil {
			qrDelay = *o.Delay
		}
		qDelay := swag.FormatInt64(qrDelay)
		if qDelay != "" {

			if err := r.SetQueryParam("delay", qDelay); err != nil {
				return err
			}
		}
	}

	if o.Device != nil {

		// query param device
		var qrDevice string

		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {

			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}
	}

	if o.Element != nil {

		// query param element
		var qrElement string

		if o.Element != nil {
			qrElement = *o.Element
		}
		qElement := qrElement
		if qElement != "" {

			if err := r.SetQueryParam("element", qElement); err != nil {
				return err
			}
		}
	}

	if o.FileType != nil {

		// query param fileType
		var qrFileType string

		if o.FileType != nil {
			qrFileType = *o.FileType
		}
		qFileType := qrFileType
		if qFileType != "" {

			if err := r.SetQueryParam("fileType", qFileType); err != nil {
				return err
			}
		}
	}

	if o.Full != nil {

		// query param full
		var qrFull bool

		if o.Full != nil {
			qrFull = *o.Full
		}
		qFull := swag.FormatBool(qrFull)
		if qFull != "" {

			if err := r.SetQueryParam("full", qFull); err != nil {
				return err
			}
		}
	}

	// path param hash
	if err := r.SetPathParam("hash", o.Hash); err != nil {
		return err
	}

	if o.Headers != nil {

		// query param headers
		var qrHeaders string

		if o.Headers != nil {
			qrHeaders = *o.Headers
		}
		qHeaders := qrHeaders
		if qHeaders != "" {

			if err := r.SetQueryParam("headers", qHeaders); err != nil {
				return err
			}
		}
	}

	if o.Height != nil {

		// query param height
		var qrHeight int64

		if o.Height != nil {
			qrHeight = *o.Height
		}
		qHeight := swag.FormatInt64(qrHeight)
		if qHeight != "" {

			if err := r.SetQueryParam("height", qHeight); err != nil {
				return err
			}
		}
	}

	if o.HideCookieBanners != nil {

		// query param hide_cookie_banners
		var qrHideCookieBanners bool

		if o.HideCookieBanners != nil {
			qrHideCookieBanners = *o.HideCookieBanners
		}
		qHideCookieBanners := swag.FormatBool(qrHideCookieBanners)
		if qHideCookieBanners != "" {

			if err := r.SetQueryParam("hide_cookie_banners", qHideCookieBanners); err != nil {
				return err
			}
		}
	}

	if o.Invalidate != nil {

		// query param invalidate
		var qrInvalidate bool

		if o.Invalidate != nil {
			qrInvalidate = *o.Invalidate
		}
		qInvalidate := swag.FormatBool(qrInvalidate)
		if qInvalidate != "" {

			if err := r.SetQueryParam("invalidate", qInvalidate); err != nil {
				return err
			}
		}
	}

	if o.Js != nil {

		// query param js
		var qrJs string

		if o.Js != nil {
			qrJs = *o.Js
		}
		qJs := qrJs
		if qJs != "" {

			if err := r.SetQueryParam("js", qJs); err != nil {
				return err
			}
		}
	}

	if o.Language != nil {

		// query param language
		var qrLanguage string

		if o.Language != nil {
			qrLanguage = *o.Language
		}
		qLanguage := qrLanguage
		if qLanguage != "" {

			if err := r.SetQueryParam("language", qLanguage); err != nil {
				return err
			}
		}
	}

	if o.Latitude != nil {

		// query param latitude
		var qrLatitude float64

		if o.Latitude != nil {
			qrLatitude = *o.Latitude
		}
		qLatitude := swag.FormatFloat64(qrLatitude)
		if qLatitude != "" {

			if err := r.SetQueryParam("latitude", qLatitude); err != nil {
				return err
			}
		}
	}

	if o.LazyloadScroll != nil {

		// query param lazyload_scroll
		var qrLazyloadScroll bool

		if o.LazyloadScroll != nil {
			qrLazyloadScroll = *o.LazyloadScroll
		}
		qLazyloadScroll := swag.FormatBool(qrLazyloadScroll)
		if qLazyloadScroll != "" {

			if err := r.SetQueryParam("lazyload_scroll", qLazyloadScroll); err != nil {
				return err
			}
		}
	}

	if o.Longitude != nil {

		// query param longitude
		var qrLongitude float64

		if o.Longitude != nil {
			qrLongitude = *o.Longitude
		}
		qLongitude := swag.FormatFloat64(qrLongitude)
		if qLongitude != "" {

			if err := r.SetQueryParam("longitude", qLongitude); err != nil {
				return err
			}
		}
	}

	if o.Proxy != nil {

		// query param proxy
		var qrProxy string

		if o.Proxy != nil {
			qrProxy = *o.Proxy
		}
		qProxy := qrProxy
		if qProxy != "" {

			if err := r.SetQueryParam("proxy", qProxy); err != nil {
				return err
			}
		}
	}

	if o.Quality != nil {

		// query param quality
		var qrQuality int64

		if o.Quality != nil {
			qrQuality = *o.Quality
		}
		qQuality := swag.FormatInt64(qrQuality)
		if qQuality != "" {

			if err := r.SetQueryParam("quality", qQuality); err != nil {
				return err
			}
		}
	}

	if o.RandomUserAgent != nil {

		// query param random_user_agent
		var qrRandomUserAgent bool

		if o.RandomUserAgent != nil {
			qrRandomUserAgent = *o.RandomUserAgent
		}
		qRandomUserAgent := swag.FormatBool(qrRandomUserAgent)
		if qRandomUserAgent != "" {

			if err := r.SetQueryParam("random_user_agent", qRandomUserAgent); err != nil {
				return err
			}
		}
	}

	if o.Redirect != nil {

		// query param redirect
		var qrRedirect bool

		if o.Redirect != nil {
			qrRedirect = *o.Redirect
		}
		qRedirect := swag.FormatBool(qrRedirect)
		if qRedirect != "" {

			if err := r.SetQueryParam("redirect", qRedirect); err != nil {
				return err
			}
		}
	}

	if o.Scale != nil {

		// query param scale
		var qrScale float64

		if o.Scale != nil {
			qrScale = *o.Scale
		}
		qScale := swag.FormatFloat64(qrScale)
		if qScale != "" {

			if err := r.SetQueryParam("scale", qScale); err != nil {
				return err
			}
		}
	}

	if o.Timezone != nil {

		// query param timezone
		var qrTimezone string

		if o.Timezone != nil {
			qrTimezone = *o.Timezone
		}
		qTimezone := qrTimezone
		if qTimezone != "" {

			if err := r.SetQueryParam("timezone", qTimezone); err != nil {
				return err
			}
		}
	}

	// path param token
	if err := r.SetPathParam("token", o.Token); err != nil {
		return err
	}

	if o.TTL != nil {

		// query param ttl
		var qrTTL int64

		if o.TTL != nil {
			qrTTL = *o.TTL
		}
		qTTL := swag.FormatInt64(qrTTL)
		if qTTL != "" {

			if err := r.SetQueryParam("ttl", qTTL); err != nil {
				return err
			}
		}
	}

	// query param url
	qrURL := o.URL
	qURL := qrURL
	if qURL != "" {

		if err := r.SetQueryParam("url", qURL); err != nil {
			return err
		}
	}

	if o.UserAgent != nil {

		// query param user_agent
		var qrUserAgent string

		if o.UserAgent != nil {
			qrUserAgent = *o.UserAgent
		}
		qUserAgent := qrUserAgent
		if qUserAgent != "" {

			if err := r.SetQueryParam("user_agent", qUserAgent); err != nil {
				return err
			}
		}
	}

	if o.Wait != nil {

		// query param wait
		var qrWait string

		if o.Wait != nil {
			qrWait = *o.Wait
		}
		qWait := qrWait
		if qWait != "" {

			if err := r.SetQueryParam("wait", qWait); err != nil {
				return err
			}
		}
	}

	if o.Width != nil {

		// query param width
		var qrWidth int64

		if o.Width != nil {
			qrWidth = *o.Width
		}
		qWidth := swag.FormatInt64(qrWidth)
		if qWidth != "" {

			if err := r.SetQueryParam("width", qWidth); err != nil {
				return err
			}
		}
	}

	if o.X != nil {

		// query param x
		var qrX int64

		if o.X != nil {
			qrX = *o.X
		}
		qX := swag.FormatInt64(qrX)
		if qX != "" {

			if err := r.SetQueryParam("x", qX); err != nil {
				return err
			}
		}
	}

	if o.Y != nil {

		// query param y
		var qrY int64

		if o.Y != nil {
			qrY = *o.Y
		}
		qY := swag.FormatInt64(qrY)
		if qY != "" {

			if err := r.SetQueryParam("y", qY); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}