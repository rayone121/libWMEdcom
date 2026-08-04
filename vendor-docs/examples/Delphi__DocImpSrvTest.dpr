program DocImpSrvTest;

uses
  Forms,
  UTest in 'UTest.pas' {TestDocImp};

{$R *.RES}

begin
  Application.Initialize;
  Application.CreateForm(TTestDocImp, TestDocImp);
  Application.Run;
end.
