Function Get-Picks {

    param(
        [Parameter(Mandatory = $true)]$filePath,
        [Parameter(Mandatory = $true)]$pickList,
        [Parameter(Mandatory = $true)]$finalDir
    )

    if (!(Test-Path $finalDir)) {
        New-Item -Path $finalDir -ItemType Directory
    }

    $picklist = Get-Content $pickList

    foreach ($pick in $pickList) {
        $pick = $pick.trim()
        $found = $false
        foreach ($photo in (Get-ChildItem $filePath)) {

            if ($photo.Name -like "*$pick*") {
                Copy-Item -Path "$filePath\$photo" -Destination $finalDir
                $found = $true
            }
            if ($found) {
                break
            }

        }

        if (!$found) {
            Write-Host "Could not find $pick :( "
        }

    }
    
}