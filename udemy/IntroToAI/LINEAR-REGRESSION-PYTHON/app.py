import argparse
import logging
import os
import sys

import numpy as np
import pandas as pd

CONFIG = {
    "default_csv": "house_data.csv",
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

    # handle outliers

    # ensure numeric data types for modeling


def main():
    # parsing command line arguments
    args = parse_arguments()

    # load and preprocess the data
    df = load_data(args.file)
    processed_df = preprocess_data(df)

    # prepare data for modeling

    # split data into traning and testing sets

    # train a model

    # evaluate the model on both training and testing sets

    # print results

    # create a visualization

    # predict price for houses not in our dataset


if __name__ == "__main__":
    main()
