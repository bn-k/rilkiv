
# Start

Copy `.env.template` to `.env` and fill gmail credential.

Run `make start`



## Routes

<details>
<summary>`/register`</summary>

- [RequestID]()
- [Logger]()
- [Recoverer]()
- [URLFormat]()
- [SetContentType.func1]()
- **/register**
  - _POST_
  - [(*Handlers).Register.func1]()
    
Body:
```json
{
        "email": "your@email.com",
        "password": "password",
        "firstname": "firstnam",
        "lastname": "lastname"
}
```

</details>

<details>
<summary>`/confirm`</summary>

- [RequestID]()
- [Logger]()
- [Recoverer]()
- [URLFormat]()
- [SetContentType.func1]()
- **/confirm**
  - _POST_
  - [(*Handlers).Confirm.func1]()

</details>
<details>

<summary>`/login`</summary>

- [RequestID]()
- [Logger]()
- [Recoverer]()
- [URLFormat]()
- [SetContentType.func1]()
- **/login**
  - _POST_
  - [(*Handlers).Login.func1]()
    
Body:
```json
{
    "email": "your@email.com",
    "password": "password"
}
```

</details>

Authentication header:

Key: `Authorization`

Value: `BEARER {token}`


<details>
<summary>`/transaction`</summary>

- [RequestID]()
- [Logger]()
- [Recoverer]()
- [URLFormat]()
- [SetContentType.func1]()
- **/transaction**
  - _POST_
  - [v5.Verifier.func1]()
  - [v5.Authenticator]()
  - [(*Handlers).MakeTransaction.func1]()
    
```json
{
    "from_wallet_id": "84012789-0ec0-4640-be96-315da927885b",
    "to_address": "69cq25vwpbykesj9qg6tc8bkwuqpnm38",
    "amount": 2
}
```

</details>

<details>
<summary>`/transactions`</summary>

- [RequestID]()
- [Logger]()
- [Recoverer]()
- [URLFormat]()
- [SetContentType.func1]()
- **/transactions**
  - _GET_
  - [v5.Verifier.func1]()
  - [v5.Authenticator]()
  - [(*Handlers).GetTransactions.func1]()
  
- Query parameters
  - `address` select address to list transactions (optional)
  - `from` and `to`: select time, fmt `2019-10-12` (optional)
  - `page` : select the page to fetch out, default is `1`
  
>  localhost:8080/transactions?address=30buIC9uxkKAafntWm6aa8R0mql5CqjfH&from=2000-01-01&to=2020-12-12&page=2

</details>

<details>
<summary>`/wallets`</summary>

- [RequestID]()
- [Logger]()
- [Recoverer]()
- [URLFormat]()
- [SetContentType.func1]()
- **/wallets**
  - _GET_
  - [v5.Verifier.func1]()
  - [v5.Authenticator]()
  - [(*Handlers).GetWallets.func1]()

</details>
