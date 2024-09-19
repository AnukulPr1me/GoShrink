Package routes

import (

)

type request struct{
	URL				string				`json:"url"`
	CustomShort		string				`json:"short"`
	Expiry			time.Duration		`json:"expiry"`
}

type response struct{
	URL					string 			`json"url"`
	CustomShort			string			`json"short`
	Expiry				time.Duration	`json:"expiry"`
	XRateRemaining		int				`json:"rate_limit"`
	XRateLimitRest		time.Duration	`json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.ctx) error {
	body:=new(request)
	if err:=c.BodyParser(&body); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
	}


	if !govalidator.IsURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid URL",
        })
	}

	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.statusServiceUnavailable).JSON(fiber.Map{
            "error": "URL contains disallowed domain",
        })
	}

	body.URL = helpers.EnforcHTTP(body.URL)
}