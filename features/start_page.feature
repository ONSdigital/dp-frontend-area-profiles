Feature: Groups

    Scenario: GET /areas and checking the response status 200
    When I GET "/areas"
    Then the HTTP status code should be "200"