Function Compare-PlexDB {
    
    param(

        [Parameter(Mandatory = $true)][String]$csvDB1,
        [Parameter(Mandatory = $true)][String]$csvDB2

    )

    $csv1 = Import-Csv $csvDB1
    $csv2 = Import-Csv $csvDB2
    $need = @()
    Remove-Variable -Name csvdb1, csvdb2

    $hash1 = @{}
    $hash2 = @{}

    foreach ($entry in $csv1) {
        if (!$hash1.ContainsKey($entry."Sort title")) {
            $hash1.add($entry."Sort title", $entry."Part Size as Bytes")
        }
    }
    foreach ($entry in $csv2) {
        if (!$hash2.ContainsKey($entry."Sort title")) {
            $hash2.add($entry."Sort Title", $entry."Part Size as Bytes")
        }
    }
    
    foreach ($movie in $csv1) {
        if (!$hash2.ContainsKey($movie."Sort title")) {
            $need += $movie
        }
    }
    
    $need | Write-Output
    #$csv1 | 
    #Select-Object "Sort title","Year","Bitrate","Video Resolution","Video Codec","Part Size","Part File","Media ID","Title","Duration","Added","Updated","Width","Height","Aspect Ratio","Audio Codec","Container","Part File Combined","Part File Path"

    #Remove-Variable -Name csv1, csv2, hash1, hash2

}

Clear-Host

Measure-Command{
Compare-PlexDB -csvDB1 "C:\Strom\Movies.fazz.1 - Movies.fazz.1.csv" -csvDB2 "C:\Strom\MoviesLvl6Mar_23_2021.csv" | Sort-Object -Unique -Property "Sort title" |  Export-Csv C:/strom/test3.csv -NoTypeInformation
} | Select-Object TotalSeconds
#Compare-PlexDB -csvDB2 "C:\Strom\Movies.fazz.1 - Movies.fazz.1.csv" -csvDB1 "C:\Strom\MoviesLvl6Mar_23_2021.csv" | Export-Csv C:/strom/test2.csv
