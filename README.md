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

But how do we work for complex configuration 
that consists of arrays of objects? For example:
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

It is difficult to split this configuration to several environment variables.

This package is aimed to deal with such configs.

Just put the config into single environment variable.
It could be multiline, but work with single line is simplier.

Running program example:

```bash

CONFIG='{"nodes":[{"url":"http:\/\/1.localhost\/","priority":1,"enabled":true},{"url":"http:\/\/2.localhost\/","priority":2,"enabled":false}]}' ./app
```

Program code example:
```GO
import jsonEnvGo "github.com/antelman107/json-env-go"

type Config  struct {
	Nodes []struct {
		URL      string `json:"url"`
		Priority int    `json:"priority"`
		Enabled  bool   `json:"enabled"`
	} `json:"nodes"`
}

func main() {
      var cfg Config
      if err := jsonEnvGo.DecodeConfigFromEnv(envName, &cfg); err != nil {
            logger.Error(err)
            return
      }
      
      // cfg now filled, do something with it
      // ...
}
```


