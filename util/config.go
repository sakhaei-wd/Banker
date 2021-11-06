package util

import "github.com/spf13/viper"

//in order to get the value of the variables and store them in this struct,
//we need to use the unmarshaling feature of Viper.
type Config struct{
	DBDriver      string `mapstructure:"DB_DRIVER"` //based on exact name of variable in app.env
    DBSource      string `mapstructure:"DB_SOURCE"`
    ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
    viper.AddConfigPath(path) 		//tell Viper the location of the config file.In this case, the location is given by the input path argument
    viper.SetConfigName("app")		//tell Viper to look for a config file with a specific name. Our config file is app.env, so its name is app
    viper.SetConfigType("env")		//tell Viper the type of the config file, which is env 
	
    //we also want viper to read values from environment variables. So :
	//tell viper to automatically override values that it has read from config file with the 
	//values of the corresponding environment variables if they exist
	viper.AutomaticEnv()

	//start reading config values. 
	 err = viper.ReadInConfig()
    if err != nil {
        return
    }
	
	//unmarshals the values into the target config object
    err = viper.Unmarshal(&config)
    return

}

