Feature: Groups

    Scenario: GET /areas and checking the response status 200
    When I navigate to "/areas"
    Then the beta phase banner should be visible
    And the improve this page banner should be visible
    And the page should have the following content
        """
            {
                "h1.ons-u-fs-xxxl": "Find facts and figures about areas in England or Wales"
            }
        """
    And should match snapshot "start_page"