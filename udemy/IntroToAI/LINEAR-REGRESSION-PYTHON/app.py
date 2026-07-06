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

def main():
    # parsing command line arguments

    # load and preprocess the data

if __name__ == "__main__":
    main()
