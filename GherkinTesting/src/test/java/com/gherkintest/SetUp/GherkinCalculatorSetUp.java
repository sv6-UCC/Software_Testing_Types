package com.gherkintest.SetUp;

import com.gherkintest.GherkinCalculator;
import static org.junit.Assert.assertEquals;
import io.cucumber.java.en.Given;
import io.cucumber.java.en.When;
import io.cucumber.java.en.Then;

public class GherkinCalculatorSetUp {

    public double result;
    public double numberOne;
    public double numberTwo;

    @Given("^The first number is (.+)$")
    public void number_one_is(double numberOne) throws Throwable {
        this.numberOne = numberOne;
    }

    @Given("^The second number is (.+)$")
    public void number_two_is(double numberTwo) throws Throwable {
        this.numberTwo = numberTwo;
    }

    @When("^The two numbers are subtracted$")
    public void subtracting() throws Throwable {
        GherkinCalculator calculator = new GherkinCalculator();
        result = calculator.subtraction(numberOne, numberTwo);
    }
    @When("^The two numbers are added$")
    public void adding() throws Throwable {
        GherkinCalculator calculator = new GherkinCalculator();
        result = calculator.addition(numberOne, numberTwo);
    }

    @When("^The two numbers are divided$")
    public void dividing() throws Throwable {
        GherkinCalculator calculator = new GherkinCalculator();
        result = calculator.division(numberOne, numberTwo);
    }
    @When("^The two numbers are multiplied$")
    public void multiplying() throws Throwable {
        GherkinCalculator calculator = new GherkinCalculator();
        result = calculator.multiplication(numberOne, numberTwo);
    }

    @Then("^The result should be (.+)$")
    public void final_result(double result) throws Throwable {
        assertEquals(this.result, result, 0.0001);
    }
}
