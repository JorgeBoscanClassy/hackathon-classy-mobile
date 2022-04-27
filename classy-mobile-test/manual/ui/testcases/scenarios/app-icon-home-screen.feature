#  ******************************************************************************
#             Author:  Emad Borjian
#       Date Created:  04-27-2022
#        Description:  This feature file is intended to test scenarios related to 
#                      how the app icon is rendered on the Home Screen.
#  ******************************************************************************

Feature: Display of the app icon on the Home Screen

Background:
  # DEV ENVIRONMENT
  * def classyMobileApp = 'exp://34-hbi.classyomid.classy-mobile.exp.direct:80'
  # END

  # STAGING ENVIRONMENT 
  # END

  # PROD ENVIRONMENT
  # END

  # GENERAL
  # END

@P3
@Regression
@TestCaseKey=TE-TC-10001
@Route=exp://34-hbi.classyomid.classy-mobile.exp.direct:80
Scenario: Verify the app icon is rendered correctly on the Home Screen

  # Ensure app is installed
  Given an admin installs classyMobileApp through 'App Store' on an iOS device
  
  # Navigate to app icon
  When navigate to 'Home Screen' page that app icon is located
  
  # Verify app icon is rendered correctly
  Then verify app icon contains Classy logo on coral background