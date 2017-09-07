*** Settings ***
Library    	  Telnet

*** Variables ***

*** Test cases ***
Set GPS to Sheffield, wait 10 secs, then to London, wait 10 secs, then Birmingham
    Set GPS Location

*** Keywords ***
# NOTE: Requires app MockGeoFix by Lukas Vacek
Set GPS Location
    Open Connection   192.168.1.28    port=5554    prompt=OK
    Execute Command    geo fix -1.4701 53.3811     # Sets location to Sheffield
    Sleep    10s
    Execute Command    geo fix -0.1278 51.5074     # Sets location to London
    Sleep    10s
    Execute Command    geo fix -1.8904 52.4862     # Sets location to Birmingham
    Close Connection
