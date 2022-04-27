#  ******************************************************************************
#             Author:  Emad Borjian
#       Date Created:  04-27-2022
#        Description:  This feature file is intended to test scenarios related to 
#                      how the app icon is rendered in the App Store.
#  ******************************************************************************

Feature: Display of the app icon in the App Store

Background:
  # DEV ENVIRONMENT
  * def classyMobileApp = 'exp://34-hbi.classyomid.classy-mobile.exp.direct:80'
  # END

  # STAGING ENVIRONMENT 
  # END

  # PROD ENVIRONMENT
  # END

  # GENERAL
  * def appName = 'Classy Mobile'
  # END

@P3
@Regression
@TestCaseKey=TE-TC-10002
@Route=exp://34-hbi.classyomid.classy-mobile.exp.direct:80
Scenario: Verify the app icon is rendered correctly in the App Store

  # Launch App Store
  Given an admin launches 'App Store' on an iOS device
  
  # Locate the app
  When select Search tab within tab bar on bottom
  And input appName in 'Games, Apps, Stories, and More' field on top
  And select 'search' button
  
  # Verify app is found, and its icon is rendered correctly
  Then verify app called appName is returned as a result
  And verify app icon contains Classy logo on coral background