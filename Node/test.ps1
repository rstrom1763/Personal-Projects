

    Invoke-WebRequest -Uri Http://192.168.0.190:8081/write -method Post -Body '{"ComputerName":"PRPRL-TESTCOMPUTER","TEST":"BRO","Test2":"bro2"}' -ContentType "application/json"
