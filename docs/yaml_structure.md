# YAML structure for commands

## GET

```yaml
command:get
spec:
  ip:<ip>
  port:<port>
```

## POST

```yaml
command:post
spec:
  - url:<url>
  data:<json>
```

## PUT

```yaml
command:put
spec:
  - url:<url>
  data:<json>
```

## PATCH

```yaml
command:patch
spec:
  - url:<url>
  data:<json>
```

## DELETE

```yaml
command:delete
spec:
  - url:<url>
```

## Notes

- use whitespaces over tabs for easy parsing -> 2 whitespace indentation
