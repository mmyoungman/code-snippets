*** Settings ***
Library           AppiumLibrary
Test Teardown    Close Application

*** Variables ***

*** Test cases ***
#Test App Opens And Finds Text "Sheffield"
#    Set Up And Open Android Application
#    Wait Until Page Contains    Sheffield    5s

#Click "updatePosition" Element
#    Set Up And Open Android Application
#    #Click Element   accessibility_id=updatePosition # accessibility_id refers to automationText in app/views/main/main-page.xml
#    Click Element   id=updatePosition  # id to id in same file
#    Wait Until Page Contains    Sheffield    5s

Click showMap Element
    Set Up And Open Android Application
    Click Element   accessibility_id=showMap
    Click A Point   10    10
    Go Back
    Wait Until Page Contains    Sheffield    5s

*** Keywords ***
Set Up And Open Android Application
    Open Application    http://localhost:4723/wd/hub    platformName=Android   deviceName=192.168.1.24    app=${CURDIR}${/}helloworld-release-mmy-recompile.apk  automationName=appium    appPackage=org.nativescript.helloworld
    # Wait Until Page Contains    Waiting    5s
