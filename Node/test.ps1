
Measure-Command{
foreach ($num in 1..1000){
    
    #Start-Sleep -Milliseconds 2
    invoke-webrequest -uri Http://3.140.244.63:8081 -method post -body (Get-Random -Minimum 11111111)

}
}