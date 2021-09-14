Function Compare-PlexDB {
    
    param(

        [Parameter(Mandatory = $true)][String]$csvDB1,
        [Parameter(Mandatory = $true)][String]$csvDB2,
        [string]$excludeFile

    )

    $csv1 = Import-Csv $csvDB1
    $csv2 = Import-Csv $csvDB2
    $need = @()
    Remove-Variable -Name csvdb1, csvdb2

    $hash = @{}
    $excludeHash = @{}
    $i = 1
    [int64]$spaceNeeded = 0

    if ($null -ne $excludeFile) {
        if (Test-Path $excludeFile) {
            foreach ($entry in (Get-Content $excludeFile)) {
                $excludeHash.add(($entry), $i)
                $i++
            }
        }
        else {
            Write-Host "$excludeFile does not exist."
        }
    }
    foreach ($entry in $csv2) {
        if (!$hash.ContainsKey($entry."Sort title")) {
            $hash.add(($entry."Sort Title"), $entry."Part Size as Bytes")
        }
    }
    foreach ($movie in $csv1) {
        if (!$hash.ContainsKey($movie."Sort title") -and !$excludeHash.ContainsKey($movie."Sort title")) {
            $need += $movie
            if (!($movie."Part Size as Byte" -like "*-*")) {
                try {
                    $spaceNeeded += [int64]$movie."Part Size as Bytes"
                }
                catch {
                    Write-Host $_
                    Write-Host $movie."Sort Title"
                }
            }
        }
    }
    $spaceNeeded = [math]::Round(($spaceNeeded / 1GB))
    Write-Host ("Space needed: $spaceNeeded GB")
    $need | Write-Output
    #$csv1 | 
    #Select-Object "Sort title","Year","Bitrate","Video Resolution","Video Codec","Part Size","Part File","Media ID","Title","Duration","Added","Updated","Width","Height","Aspect Ratio","Audio Codec","Container","Part File Combined","Part File Path"

    #Remove-Variable -Name csv1, csv2, hash1, hash

}

#Clear-Host

Compare-PlexDB -csvDB1 "C:\Strom\kentLibrary.csv" -csvDB2 "C:/strom/RyanLibrary.csv" -excludeFile C:/strom/exclude.txt | Sort-Object -Unique -Property "Sort title" |  Export-Csv C:/strom/KentNoHave.csv -NoTypeInformation

#Compare-PlexDB -csvDB1 "C:/strom/RyanLibrary.csv" -csvDB2 "C:\Strom\kentLibrary.csv" | Sort-Object -Unique -Property "Sort title" | Export-Csv C:/strom/RyanNoHave.csv -NoTypeInformation
