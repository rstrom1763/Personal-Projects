$picks = Get-Content C:/strom/picks.txt

foreach ($jpg in (Get-ChildItem 'C:/strom/photos/chief_mann_retirement/jpg')) {
    foreach ($pick in $picks) {
        if ($jpg.name -like "*$pick*") {
            Copy-Item -LiteralPath $jpg.Name -Destination 'C:/strom/photos/chief_mann_retirement/picks/jpg/'
        }
    }
}