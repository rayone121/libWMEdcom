# libWMEdcom

Go library for the WinMENTOR DocImpServer DCOM interface. Wraps all 145 COM methods with a goroutine-safe client.

## Requirements

- Windows with WinMENTOR installed (DocImpServer.dll registered)
- Go 1.22+

## Installation

```bash
go get github.com/rayone121/libWMEdcom
```

## Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/rayone121/libWMEdcom/winmentor"
)

func main() {
    client, err := winmentor.NewClient()
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // Set active company and work month
    client.SetNumeFirma("MYCOMPANY")
    client.SetLunaLucru(2024, 10)

    // Query data
    partners, _ := client.GetListaParteneri()
    fmt.Printf("Found %d partners\n", len(partners))

    stock, _ := client.GetStocArticole()
    fmt.Printf("Found %d stock articles\n", len(stock))

    // Import invoices
    lines := []string{"header;...", "line1;...", "line2;..."}
    client.SetDocsData(lines)
    if err := client.DateValide(); err != nil {
        log.Fatal("Validation failed:", err)
    }
    count, _ := client.ImportaFacturi()
    fmt.Printf("Imported %d invoices\n", count)
}
```

## Architecture

All COM calls are dispatched to a dedicated goroutine locked to an OS thread. The client is safe to use from any goroutine.

### Packages

- `winmentor` - Main client package

### Method Categories

| Category | Methods | Description |
|----------|---------|-------------|
| Session | `SetNumeFirma`, `SetLunaLucru`, `GetListaFirme`, `GetListaLuni` | Company/month selection |
| Partners | `GetListaParteneri`, `AdaugaPartener`, `ModificaPartener`, `GetSoldPart`, `GetSolduriExt` | Partner CRUD and balances |
| Articles | `GetStocArticole`, `GetNomenclatorArticole`, `GetProducts`, `AddProduct` | Stock, nomenclature, products |
| Documents | `SetDocsData` → `DateValide` → `ImportaFacturi` | Invoice/order import workflow |
| Orders | `ComenziValide`, `ImportaComenzi`, `GetComenziNefacturate` | Order management |
| Cash Register | `MonetareValide`, `ImportaMonetare` | Cash register operations |
| Queries | `GetVanzariExt`, `GetListaPersonal`, `GetListaGestiuni`, `GetListaBanci` | Sales, employees, warehouses |
| Config | `SetIDPartField`, `SetIDArtField`, `SetDescarcareAutomata` | Server configuration |

### Document Import Workflow

1. `SetDocsData(lines)` - Send document data
2. `DateValide()` / `ComenziValide()` / etc. - Validate
3. `ImportaFacturi()` / `ImportaComenzi()` / etc. - Import
4. `GetListaErori()` - Check errors if validation fails

## License

MIT
