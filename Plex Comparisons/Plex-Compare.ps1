Function Compare-PlexDB {
    
    param(

        [Parameter(Mandatory = $true)][String]$csvDB1,
        [Parameter(Mandatory = $true)][String]$csvDB2

    )

    $need = @()

    $csv1 = Import-Csv $csvDB1
    $csv2 = Import-Csv $csvDB2

    Remove-Variable -Name csvdb1, csvdb2

    #$csv1 | 
    #Select-Object "Sort title","Year","Bitrate","Video Resolution","Video Codec","Part Size","Part File","Media ID","Title","Duration","Added","Updated","Width","Height","Aspect Ratio","Audio Codec","Container","Part File Combined","Part File Path"

    $count = 0
    foreach ($movie in $csv1) {
        
        $found = $false

        foreach ($movie2 in $csv2) {

            if ($movie."Sort title" -eq $movie2."Sort title") {

                $found = $true
            
            }
            
            if ($found) {
                
                break

            }
        
        }
        if (!$found) {
            
            $need += $movie
            $count += 1
        
        }

    }

    $need | Sort-Object -Property "Sort title" | Write-Output
    Remove-Variable -Name need, count, csv1, csv2

}

Clear-Host
measure-command {Compare-PlexDB -csvDB1 "C:\Strom\Movies.fazz.1 - Movies.fazz.1.csv" -csvDB2 "C:\Strom\MoviesLvl6Mar_23_2021.csv" | Sort-Object -Unique -Property "Sort title" | Export-Csv C:/strom/test.csv -NoTypeInformation } | Select-Object TotalSeconds
#Compare-PlexDB -csvDB2 "C:\Strom\Movies.fazz.1 - Movies.fazz.1.csv" -csvDB1 "C:\Strom\MoviesLvl6Mar_23_2021.csv" | Export-Csv C:/strom/test2.csv
