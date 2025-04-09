Feature: calculator for Gherkin testing

    Scenario Outline: Add two numbers
        Given The first number is <firstNumber>
        And The second number is <secondNumber>
        When The two numbers are added
        Then The result should be <result>

        Examples:
            | firstNumber | secondNumber | result |
            | 50    | 70     | 120    |
            | 20    | 60     | 80     |
            | 30    | 40     | 70     |

    Scenario Outline: Subtract two numbers
        Given The first number is <firstNumber>
        And The second number is <secondNumber>
        When The two numbers are subtracted
        Then The result should be <result>

        Examples:
            | firstNumber | secondNumber | result |
            | 50    | 70     | -20    |
            | 10    | 20     | -10    |
            | 30    | 40     | -10    |

    Scenario Outline: Multiply two numbers
        Given The first number is <firstNumber>
        And The second number is <secondNumber>
        When The two numbers are multiplied
        Then The result should be <result>

        Examples:
            | firstNumber | secondNumber | result |
            | 50    | 70     | 3500   |
            | 10    | 20     | 200    |
            | 30    | 40     | 1200   |

    Scenario Outline: Divide two numbers
        Given The first number is <firstNumber>
        And The second number is <secondNumber>
        When The two numbers are divided
        Then The result should be <result>

        Examples:
            | firstNumber | secondNumber | result |
            | 50    | 70     | 0.7143 |
            | 10    | 20     | 0.5    |
            | 30    | 40     | 0.75   |
