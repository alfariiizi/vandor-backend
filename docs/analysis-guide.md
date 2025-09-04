input_products (array)

```
[
  {"product_id":123, "source":"tokopedia", "source_id":"98765", "url":"https://..."},
  {"product_id":124, "source":"shopee", "source_id":"5555", "url":"https://..."}
]
```

top_products (array, ordered)

```
[
  {
    "rank": 1,
    "score": 0.93,
    "product_ref": {"product_id": 123, "source":"tokopedia", "source_id":"98765", "url":"..."},
    "short_reason": "High margin, low competition in subcategory; stable price over time"
  },
  {
    "rank": 2,
    "score": 0.87,
    "product_ref": {"product_id": 124, "source":"shopee", "source_id":"5555", "url":"..."},
    "short_reason": "Good reviews and consistent sales but lower margin"
  }
]
```

winner

```
{
  "product_ref": {"product_id": 123, "source":"tokopedia"},
  "score": 0.93,
  "short_reason": "Best combination of margin, demand and low competition"
}
```

analysis_texts

```
{
  "top_10_reasoning": "Long — very-very detailed plain text explaining each top product...",
  "winner_reasoning": "Very detailed explanation why product X is best — includes data points and assumptions...",
  "suggestion": "Step-by-step recommendation on what product to build, how to manufacture, channels to sell, go-to-market plan, KPIs..."
}
```

budget example

```
{
  "production_cost": {"currency":"IDR","unit_cost":75000,"units_first_run":1000,"total":75000000},
  "marketing_budget": {"currency":"IDR","amount":15000000,"channels": {"tiktok":0.6,"facebook":0.3,"google":0.1}},
  "assumptions":"manufacturing in China, MOQ 1000, shipping CIF Jakarta"
}
```
