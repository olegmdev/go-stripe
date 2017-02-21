package main

import "fmt"
import "wrapper/payments"
import "github.com/spf13/viper"

func init() {
 viper.SetConfigType("yaml")
 viper.AddConfigPath(".")
 viper.SetConfigName("config")

 err := viper.ReadInConfig()
 if err != nil {
   panic(fmt.Errorf("Fatal error config file: %s \n", err))
 }
}

func main() {
  var paymentProvider payments.IPaymentProvider
  paymentProvider = &payments.StripeProvider{
    Key: viper.Get("stripe.key").(string),
  }

  paymentMaker := PaymentMaker{ paymentProvider }

  // do a payment and have fun ;)
  paymentMaker.doAPayment()
}
