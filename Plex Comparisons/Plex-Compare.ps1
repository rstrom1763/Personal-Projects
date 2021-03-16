﻿Function Compare-PlexDB{

$csv1 = Import-Csv "C:\Users\autob\Downloads\Movies-Level 6-20210102-210814.csv"
$csv2 = Import-Csv "C:\Users\autob\Downloads\Kent-plex-library-exported - Kent-plex-library-exported.csv"
$need = @()

#$csv1 | 
#Select-Object "Sort title","Year","Bitrate","Video Resolution","Video Codec","Part Size","Part File","Media ID","Title","Duration","Added","Updated","Width","Height","Aspect Ratio","Audio Codec","Container","Part File Combined","Part File Path"

$count = 0
foreach($movie in $csv1){
    
    $found = $false

    foreach ($movie2 in $csv2){

        if($movie."Sort title" -eq $movie2."Sort title"){

            $found = $true
        
        }
        if($found){
            
            break

        }
    
    }
    if($found -eq $false){
        
        $need += $movie
        $count += 1
    
    }

}
$need | Select-Object "Sort title"
$count

}