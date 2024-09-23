package routes

import (
	"os"
	"time"
	"strconv"
	"time"
	"github.com/anukulpr1me/GoShrink/database"
	"github.com/anukulpr1me/GoShrink/helpers"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"

	
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

func ShortenURL(c *fiber.Ctx) error {
	body:=new(request)
	if err:=c.BodyParser(&body); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
	}

	r2 := database.CreateClient(1)
	defer r2.Close()
	val,err := r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil {
		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*time.Second).Err()
	}else{
		val, _ = r2.Get(database.Ctx, c.IP()).Result()
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error": "rate limit exceeded",
				"rate_limit_reset": limit/limit.Nanosecond/limit.Minute,
			})
		}
	}
	if !govalidator.IsURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid URL",
        })
	}

	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
            "error": "URL contains disallowed domain",
        })
	}

	body.URL = helpers.EnforceHTTP(body.URL)
	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	}else{
		id = body.CustomShort
	}
	r := database.CreateClient(0)
	defer r.Close()
	val, _ = r.Get(database.Ctx, id).Result()
	if(val != ""){
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "URL custom short URL is alrady exists",
		})
	}
	if body.Expiry == 0{
		body.Expiry = 24
	}

	err = r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Unable to connect to server",
        })
	}

	resp := response {
		URL: body.URL,
        CustomShort: "",
        Expiry: body.Expiry,
        XRateRemaining: 10,
        XRateLimitRest: 30,
	}
	r2.Decr(database.Ctx, c.IP())
	val, _ = r2.Get(database.Ctx, c.IP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)
	ttl, _ := r2.TTL(database.Ctx, c.IP()).Result()
	resp.XRateLimitRest = ttl/time.Nanosecond/time.Minute
	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id 
	return c.Status(fiber.StatusOK).JSON(resp)

}