# What is it?
Tiny (1 line) GO package to deal with complex configs provided through environment variables.

# Why we need it?
In dockerized applications and by https://12factor.net/ recommendation
we have to pass configuration for our GO program through environment variables.

It is easy if configuration consists of simple values (strings, numbers, boolean, etc), something like this:

```json
{
      "login":"user",
      "password":"somepass",
      "mfa_enabled": true
}
```

In this case we could use several environment variables like  `APP_LOGIN="user" APP_PASSWORD="somepass"`.

But how would we work with complex configuration that, for example, consists of arrays of objects? For example:
```json
{
      "nodes":[
            {
                  "url":"http://1.localhost/",
                  "priority":1,
                  "enabled":true
            },
            {
                  "url":"http://2.localhost/",
                  "priority":2,
                  "enabled":false
            }
      ]
}
```

This package is aimed to deal with such configs.
Just put the config into single environment variable.
It could be multiline, but work with single line is simplier from `bash`.

Running program `bash` example:

```bash

CONFIG='{"nodes":[{"url":"http:\/\/1.localhost\/","priority":1,"enabled":true},{"url":"http:\/\/2.localhost\/","priority":2,"enabled":false}]}' ./app
```

GO program code example:
```GO
import jsonEnvGo "github.com/antelman107/json-env-go"

type Config struct {
	Nodes []struct {
		URL      string `json:"url"`
		Priority int    `json:"priority"`
		Enabled  bool   `json:"enabled"`
	} `json:"nodes"`
}

func main() {
      var cfg Config
      if err := jsonEnvGo.DecodeConfigFromEnv("CONFIG", &cfg); err != nil {
            logger.Error(err)
            return
      }
      
      // cfg now filled, do something with it
      // ...
}
```

This package is very tiny, it can be used instead of viper (https://github.com/spf13/viper) when working especially with json configs coming from environment variables.


