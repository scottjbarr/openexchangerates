# Open Exchange Rates

A partial client for the Open Exchange Rates API.

Only endpoints I am interested in are implemented.


## Usage

```
appID := "yourAppID"
fx := openexchangerates.New(yourAppID)

// get exchange rates with USD as the base currency
latest, err := fx.Latest("USD")
```


## References

- https://docs.openexchangerates.org/


## Licence

The MIT License (MIT)

Copyright (c) 2017 Scott Barr

See [LICENSE.md](LICENSE.md)
