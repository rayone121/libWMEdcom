# Vendor example projects

WinMENTOR's own DocImpServer examples, from
<https://ftp.winmentor.ro/WinMentor/Documentatie/24_DocImpServer/ExempleUtilizare/>.

That host carries 154 files; `download.winmentor.ro` shows the same folder with 4. Only the
sources are kept here — the build output, `obj/` caches and generated designer files were left
behind.

## Which of these describe OUR server

**`dcom-test-paradox__*` and `C#__dcom-test-visual-paradox__*`** instantiate
`new DocImpServer.DocImpObject()`. That is our server: WinMENTOR Classic on a Paradox backend,
which is what AEG runs. Authoritative.

**`Delphi__*`** likewise. `Delphi__DocImpServer_TLB.pas` is the type library imported into Pascal
and independently confirms the dispids we extracted from the interop assembly — same numbers,
`GetListaParteneri` 7, `GetNomenclatorArticole` 25, `GetInfoIesiri` 62, `GetTransferuri` 76. Note
it was generated in 2008 and declares 89 methods against today's 145.

**Not kept, and not applicable:** the sibling `dcom-test-visual` project instantiates
`new WMDocImpServer.WMDocImpObject()`. That is WinMENTOR **ENTERPRISE** — a different product,
a different assembly, and an interface with 86 methods ours does not have, among them `LogOn`,
`GetLiniiComandaNefacturata`, `SetDenSubunit`, `SetCantitatiLiniiComanda`, `SetStadiuWMSComanda`
and `GetCmdFurnNefacturate`. The two projects sit side by side with nearly identical names and
call nearly identical methods, so check which object a file constructs before drawing any
conclusion from it.

## What is here

| File | Why |
|---|---|
| `dcom-test-paradox__dcom-test2__Program.cs` | the console harness — the fullest set of calls |
| `C#__dcom-test-visual-paradox__…Form1.cs` | the same, with a UI |
| `Delphi__UTest.pas` | the Delphi equivalent |
| `Delphi__DocImpServer_TLB.pas` | the type library as Pascal declarations |
| `dcom-test-paradox__FactIes.ini` | a real invoice packet |
| `dcom-test-paradox__partener.txt` | a real AdaugaPartener record |
| `dcom-test-paradox__test-dcom.php` | calling it from PHP |
| `*__docimpserver` | setup notes: registration, DCOM permissions, MIDAS.DLL |

## Operational notes worth knowing

Registration is `DocImpServer.exe /RegServer`, and an upgrade needs `/UnregServer` first. If a
call fails with "Error loading MIDAS.DLL", register that with `regsvr32`.

The vendor forum records that DocImpServer behaves like `Mentor.exe`: it can hold `Jurnal.db`
while it works and must be shut down before exclusive operations. That is why the web explorer
keeps one shared COM client and refuses to run while a sync is in progress.
