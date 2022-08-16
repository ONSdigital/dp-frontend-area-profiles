Feature: Groups
    Scenario: GET /areas and checking the response status 200
    When I navigate to "/areas"
    Then the beta phase banner should be visible
    And the improve this page banner should be visible
    # -- Top section 
    And the page heading should be "Facts and figures for areas in England and Wales"
    And the first paragraph should have a link of "England"
    And the first paragraph should have a second link of "Wales"
    # -- Country section
    And the country section sub heading is "Other countries"
    And the country section first paragraph contains text "Scotland and Northern Ireland have their own agencies who produce official statistics. View facts and figures for areas in:"
    And the page should have the following content
        """
            {
                "[data-test='other-countries-ul'] > li:nth-child(1) > a": "Scotland",
                "[data-test='other-countries-ul'] > li:nth-child(2) > a": "Northern Ireland"
            }       
        """
    # -- Related section
    And the Related content heading should be "Related content"
    And the page should have the following content
    """
        {
            "[data-test='related-ul'] > li:nth-child(1) > a": "Census"
        }       
    """

