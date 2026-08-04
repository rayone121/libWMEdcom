<?php 
//exec('"C:\Users\cristian.ungureanu\Documents\Visual Studio 2015\Projects\dcom-test\dcom-test\bin\Debug\dcom-test.exe"',$out);

$function="";

/*
if(isset($_GET["func"]))
	$function = htmlspecialchars($_GET["func"]);
*/

$command = "dcom-test.exe ";
foreach ($_GET as $key => $value) {
	$command .= "$value ";
}

$output = "";
//exec("dcom-test.exe $function", $output);
exec($command, $output);
print_r($output);

?>