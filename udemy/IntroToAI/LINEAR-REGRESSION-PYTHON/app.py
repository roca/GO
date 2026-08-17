"""
Housing Price Linear Regression Analysis

This script performs linear regression analysis on housing data,
predicting prices based on square footage. It includes data validation,
preprocessing, model training/testing, visualization, and the ability to
predict the price of a new house given its square footage after training.

What is Linear Regression?
-------------------------
Linear regression is a statistical method that attempts to find a linear relationship
between input variables (features) and an output variable (target).

In this case, we're trying to find the relationship between:
- Input/Feature: Square footage of a house
- Output/Target: Price of the house (in thousands of dollars)

The goal is to find the line that best fits the data points, represented by the equation:
    y = mx + b
which in this case is:
    Price = m × Square Footage + b
where:
- m is the slope (how much price increases when square footage increases by 1)
- b is the intercept (the theoretical price of a house with 0 square footage)

This "line of best fit" allows us to make predictions for new houses based on their square footage.
"""

import argparse
import logging
import os
import sys

import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_squared_error, r2_score
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import StandardScaler

CONFIG = {
    "default_csv": "house_data.csv",
    "test_size": 0.2,  # 20% of data used for testing, 80% for training
    "random_state": 42,  # Seed for random operations, ensures reproducibility
    "figure_size": (10, 6),
    "point_color": "blue",
    "line_color": "red",
    "grid_alpha": 0.3,
    "output_image": "housing_regression.png",
    "line_width": 2,
}

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] %(message)s",
)
logger = logging.getLogger(__name__)


def parse_arguments():
    parser = argparse.ArgumentParser(
        description="Linear regression analysis on housing data"
    )
    parser.add_argument(
        "-f",
        "--file",
        type=str,
        default=CONFIG["default_csv"],
        help=f"Path to the CSV file containing housing data: {CONFIG['default_csv']}",
    )
    parser.add_argument(
        "--no_plot",
        action="store_true",
        help="Do not display the plot (still saves to file)",
    )
    return parser.parse_args()


def load_data(file_path):
    if not os.path.isfile(file_path):
        logger.error(f"File not found: {file_path}")
        sys.exit(1)

    try:
        logger.info(f"Loading data from {file_path}")
        df = pd.read_csv(file_path)

        # validate for required columns
        required_columns = ["square_footage", "price_thousands"]
        for col in required_columns:
            if col not in df.columns:
                logger.error(f"Missing required column: {col}")
                sys.exit(1)

        return df

    except Exception as e:
        logger.error(f"Error loading data from {file_path}: {e}")
        sys.exit(1)


def preprocess_data(df):
    logger.info("Preprocessing data")
    processed_df = df.copy()

    # handle missing values
    if processed_df[["square_footage", "price_thousands"]].isna().any().any():
        logger.warning("Missing values found. Dropping rows with missing values.")
        processed_df = processed_df.dropna(subset=["square_footage", "price_thousands"])

    # filter out outliers
    for col in ["square_footage", "price_thousands"]:
        mean = processed_df[col].mean()
        std = processed_df[col].std()
        lower_bound = mean - 3 * std
        upper_bound = mean + 3 * std

        outliers = (processed_df[col] < lower_bound) | (processed_df[col] > upper_bound)
        if outliers.any():
            logger.warning(
                f"Outliers detected in {col}. Removing {outliers.sum()} outliers."
            )
            processed_df = processed_df[~outliers]

    # ensure numeric data types for modeling
    processed_df["square_footage"] = pd.to_numeric(
        processed_df["square_footage"], errors="coerce"
    )
    processed_df["price_thousands"] = pd.to_numeric(
        processed_df["price_thousands"], errors="coerce"
    )

    processed_df = processed_df.dropna(subset=["square_footage", "price_thousands"])

    return processed_df


def train_model(X, y):
    logger.info("Training linear regression model")
    # scale the features for better model performance
    scaler = StandardScaler()
    X_scaled = scaler.fit_transform(X)

    # train the linear regression model
    model = LinearRegression()
    model.fit(X_scaled, y)

    return model, scaler


def evaluate_model(model, X, y, scaler):
    # scale the features for evaluation
    X_scaled = scaler.transform(X)

    # make predictions
    predictions = model.predict(X_scaled)

    # calculate R-squared and RMSE
    r2 = r2_score(y, predictions)
    rmse = np.sqrt(mean_squared_error(y, predictions))

    return predictions, r2, rmse


def print_results(
    X_train,
    y_train,
    X_test,
    y_test,
    train_predictions,
    test_predictions,
    model,
    scaler,
):
    slope = model.coef_[0] / scaler.scale_[0]
    intercept = model.intercept_ - (model.coef_[0] * scaler.mean_[0] / scaler.scale_[0])

    r_squared_train = r2_score(y_train, train_predictions)
    r_squared_test = r2_score(y_test, test_predictions)
    rmse_train = np.sqrt(mean_squared_error(y_train, train_predictions))
    rmse_test = np.sqrt(mean_squared_error(y_test, test_predictions))

    print(
        f"\nLinear Regression Formula: Price = {slope:.4f} * Square Footage + {intercept:.4f}"
    )
    print(f"R-squared (train): {r_squared_train:.4f}")
    print(f"R-squared (test): {r_squared_test:.4f}")
    print(f"RMSE (train): {rmse_train:.4f}")
    print(f"RMSE (test): {rmse_test:.4f}")

    train_df = pd.DataFrame(
        {
            "Square Footage": X_train.flatten(),
            "Actual Price ($K)": y_train,
            "Predicted Price ($K)": np.round(train_predictions, 2),
        }
    )

    test_df = pd.DataFrame(
        {
            "Square Footage": X_test.flatten(),
            "Actual Price ($K)": y_test,
            "Predicted Price ($K)": np.round(test_predictions, 2),
        }
    )

    print("\nTraining Prediction Sample (first 5 rows):")
    print(train_df.head().to_string(index=False))

    print("\nTesting Prediction Sample (first 5 rows):")
    print(test_df.head().to_string(index=False))


def create_visualization(
    X_train,
    y_train,
    X_test,
    y_test,
    train_predictions,
    test_predictions,
    model,
    scaler,
    output_file,
    show_plot=True,
):
    plt.figure(figsize=CONFIG["figure_size"])

    # plot the training data
    plt.scatter(
        X_train,
        y_train,
        color=CONFIG["point_color"],
        alpha=0.7,
        label="Training data",
    )
    # plot the testing data
    plt.scatter(X_test, y_test, color="green", alpha=0.7, label="Test data")

    x_range = np.linspace(
        min(X_train.min(), X_test.min()),
        max(X_train.max(), X_test.max()),
        100,  # 100 points for a smooth line
    ).reshape(-1, 1)

    # scale the range and predict corresponding y-values
    x_range_scaled = scaler.transform(x_range)
    y_range_pred = model.predict(x_range_scaled)

    # plot the regression line
    plt.plot(
        x_range,
        y_range_pred,
        color=CONFIG["line_color"],
        linewidth=CONFIG["line_width"],
        label="Regression line",
    )

    # Add labels and title
    plt.xlabel("Square Footage")
    plt.ylabel("Price (thousands $)")
    plt.title("Linear Regression: Housing Price vc Square Footage")
    plt.legend()
    plt.grid(True, alpha=CONFIG["grid_alpha"])

    # calculate and display model parameters
    slope = model.coef_[0] / scaler.scale_[0]
    intercept = model.intercept_ - (model.coef_[0] * scaler.mean_[0] / scaler.scale_[0])
    r_squared_train = r2_score(y_train, train_predictions)
    r_squared_test = r2_score(y_test, test_predictions)

    # format text to display on the plot
    formula_text = f"Price = {slope:.4f} * Square Footage + {intercept:.4f}"
    r2_train_text = f"R² (train) = {r_squared_train:.4f}"
    r2_test_text = f"R² (test) = {r_squared_test:.4f}"

    plt.figtext(0.15, 0.85, formula_text, fontsize=12)
    plt.figtext(0.15, 0.82, r2_train_text, fontsize=12)
    plt.figtext(0.15, 0.79, r2_test_text, fontsize=12)

    plt.savefig(output_file)
    logger.info(f"Plot saved as {output_file}")

    if show_plot:
        plt.show()

    plt.close()


def main():
    # parsing command line arguments
    args = parse_arguments()

    # load and preprocess the data
    df = load_data(args.file)
    processed_df = preprocess_data(df)

    # prepare data for modeling
    X = processed_df["square_footage"].values.reshape(-1, 1)  # 2D array for sklearn
    y = processed_df["price_thousands"].values

    # split data into traning and testing sets
    X_train, X_test, y_train, y_test = train_test_split(
        X, y, test_size=CONFIG["test_size"], random_state=CONFIG["random_state"]
    )
    logger.info(
        f"Data split into training and testing sets: "
        f"{len(X_train)} training samples, {len(X_test)} testing samples"
    )

    # train a model
    model, scaler = train_model(X_train, y_train)
    logger.info("Model training completed")

    # evaluate the model on both training and testing sets
    train_predictions, train_r2, train_rmse = evaluate_model(
        model, X_train, y_train, scaler
    )
    test_predictions, test_r2, test_rmse = evaluate_model(model, X_test, y_test, scaler)

    logger.info(
        f"Model evaluation complete. R-squared (train): {train_r2:.4f}, R-squared (test): {test_r2:.4f}"
    )

    # print results
    print_results(
        X_train,
        y_train,
        X_test,
        y_test,
        train_predictions,
        test_predictions,
        model,
        scaler,
    )

    # create a visualization
    create_visualization(
        X_train,
        y_train,
        X_test,
        y_test,
        train_predictions,
        test_predictions,
        model,
        scaler,
        CONFIG["output_image"],
        not args.no_plot,
    )

    # predict price for houses not in our dataset


if __name__ == "__main__":
    main()
