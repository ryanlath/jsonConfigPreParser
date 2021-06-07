# JSON Config Pre-Parser

Golang pre-parser to allow comments (// /**/ #) and trailing commas in JSON config files.  Also, removes white space.

## Why?

I needed a config file format for golang.  I've used this, I've used that, I've rolled my own. In the end, JSON makes the most sense: it's well supported and familiar.  But no comments?  No trailing commas?  !@#$ that.  A Google search found a 10 year old [gist](https://gist.github.com/kurokikaze/1254281) that pretty much did what I wanted.  Thanks @kurokikaze!

## Example

See [test](jsonConfigPreParser_test.go) or basically
```
var config *ConfigStruct

if bytes, err := jsonConfigPreParser.Parse(bytesWithComments); err == nil {
	if err = json.Unmarshal(bytes, &config); err == nil {
		// use config
	}
}
```
## Why not YAML, TOML, INI, XYZ?

They ALL suck. :-P

## But the JSON spec says no!?!?

Whatever.

## Credits / License

Taken from [https://gist.github.com/kurokikaze/1254281](https://gist.github.com/kurokikaze/1254281).  

My two or three lines of, er, brilliance are released under the [DAL](LICENSE)!