*** Settings ***
Library           AppiumLibrary
Test Teardown    Close Application

*** Variables ***

*** Test cases ***
Test App Opens And Finds Text "Sheffield"
    Set Up And Open Android Application
    Wait Until Page Contains    Sheffield    5s

*** Keywords ***
Set Up And Open Android Application
    Open Application    http://localhost:4723/wd/hub    platformName=Android   deviceName=192.168.1.24    app=${CURDIR}${/}helloworld-release.apk  automationName=appium    appPackage=org.nativescript.helloworld
    Wait Until Page Contains    Waiting    5s
