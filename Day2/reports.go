package main

var Reports = []string{
  "6 8 9 11 14 12",
  "31 33 36 39 42 42",
  "5 6 7 9 11 13 17",
  "7 9 12 15 17 18 21 27",
  "58 59 60 59 60",
  "59 60 57 60 61 63 61",
  "68 70 71 70 72 75 76 76",
  "17 18 19 21 23 22 25 29",
  "56 57 60 58 64",
  "10 11 12 12 15",
  "87 90 93 96 97 98 98 96",
  "40 41 44 44 45 46 46",
  "72 75 75 78 79 83",
  "77 80 80 81 88",
  "43 45 48 49 50 54 56 58",
  "50 51 54 58 57",
  "44 45 47 51 51",
  "46 47 50 54 57 60 63 67",
  "12 15 17 21 24 29",
  "82 83 86 92 95",
  "18 21 26 29 27",
  "68 71 72 74 76 83 83",
  "7 10 17 20 22 26",
  "30 31 37 38 45",
  "24 21 22 24 25 26 27",
  "34 31 33 34 32",
  "4 3 5 6 8 10 13 13",
  "18 17 20 23 25 29",
  "29 27 29 30 37",
  "61 58 61 58 59 62",
  "65 62 63 60 63 65 67 65",
  "13 11 13 12 12",
  "25 23 20 23 26 29 31 35",
  "58 55 52 54 61",
  "24 22 25 25 26 27 29 31",
  "25 22 23 24 27 27 24",
  "10 8 8 9 9",
  "64 63 66 66 70",
  "72 71 74 74 77 83",
  "89 86 89 92 96 99",
  "89 87 88 92 93 96 97 95",
  "87 85 88 92 93 93",
  "77 74 78 79 82 84 88",
  "42 39 42 43 47 53",
  "79 76 79 82 88 89 90 92",
  "58 56 59 61 63 70 72 69",
  "42 40 41 46 47 47",
  "55 53 56 58 64 65 69",
  "46 43 45 50 52 57",
  "67 67 68 69 70 73 74 75",
  "61 61 63 65 62",
  "24 24 27 30 32 32",
  "30 30 31 32 33 36 40",
  "40 40 43 45 48 50 55",
  "62 62 64 66 64 66 67",
  "88 88 90 93 94 95 94 93",
  "88 88 85 88 89 90 90",
  "69 69 70 73 75 72 73 77",
  "10 10 7 10 12 17",
  "12 12 14 15 15 17 20",
  "14 14 16 17 19 19 22 20",
  "7 7 8 8 11 11",
  "89 89 92 93 94 94 98",
  "85 85 86 89 89 96",
  "70 70 74 77 79",
  "86 86 90 93 94 93",
  "26 26 29 30 32 36 36",
  "85 85 86 90 94",
  "35 35 36 40 46",
  "61 61 68 70 72 73",
  "71 71 74 76 79 80 87 86",
  "88 88 89 92 99 99",
  "22 22 28 30 32 35 39",
  "11 11 13 19 24",
  "4 8 9 11 12 15",
  "13 17 20 23 26 29 26",
  "64 68 70 71 74 74",
  "80 84 85 88 90 91 94 98",
  "71 75 76 78 81 83 86 93",
  "34 38 41 43 40 41 43 45",
  "2 6 5 8 10 9",
  "25 29 31 29 29",
  "64 68 69 72 69 73",
  "19 23 25 24 25 32",
  "65 69 69 70 71 72 75",
  "84 88 91 91 94 96 97 95",
  "86 90 93 93 93",
  "30 34 36 37 40 40 44",
  "69 73 76 78 78 79 84",
  "79 83 85 89 90 93 95",
  "4 8 11 12 16 13",
  "19 23 26 30 32 33 36 36",
  "23 27 28 29 33 34 38",
  "10 14 18 20 21 26",
  "26 30 32 38 41",
  "61 65 72 74 75 76 74",
  "35 39 41 46 47 48 48",
  "35 39 40 41 43 46 53 57",
  "57 61 68 70 75",
  "80 86 87 89 91 93",
  "59 64 67 68 71 69",
  "33 38 41 43 43",
  "55 60 61 64 65 66 67 71",
  "41 48 50 51 52 57",
  "43 50 53 56 55 57 60 63",
  "89 94 96 99 96 97 99 98",
  "11 16 18 21 23 26 23 23",
  "54 61 64 66 67 66 70",
  "28 34 35 37 36 38 40 45",
  "47 54 55 58 60 60 62 64",
  "77 84 87 89 90 90 92 90",
  "21 26 27 27 29 29",
  "50 55 58 58 60 62 64 68",
  "10 17 18 20 20 27",
  "8 14 18 19 22 23 25 27",
  "31 38 41 42 46 49 48",
  "58 64 66 70 73 76 76",
  "26 32 35 38 42 46",
  "61 68 72 75 80",
  "21 26 28 30 33 38 39",
  "30 36 38 44 46 49 47",
  "2 7 8 14 14",
  "39 44 47 49 52 59 62 66",
  "64 71 73 80 85",
  "76 75 74 72 73",
  "54 52 49 47 44 42 41 41",
  "87 84 82 81 78 76 72",
  "62 59 58 57 54 52 49 43",
  "94 92 89 86 89 88",
  "71 70 72 69 70",
  "50 49 48 51 49 46 46",
  "35 34 31 34 30",
  "20 19 18 21 18 12",
  "95 93 90 90 89 88",
  "50 47 45 44 42 42 45",
  "59 57 55 55 54 51 51",
  "54 51 51 48 46 45 41",
  "93 90 88 87 85 85 80",
  "20 17 14 10 9 7 4",
  "90 89 88 85 84 80 81",
  "43 41 37 36 33 33",
  "97 94 90 87 86 82",
  "15 13 9 8 2",
  "88 86 84 82 75 74 72 69",
  "64 63 58 56 54 57",
  "12 10 5 3 2 2",
  "65 63 61 60 59 54 50",
  "44 41 36 33 28",
  "80 82 81 80 77 76 73 72",
  "78 81 79 77 78",
  "21 24 23 22 21 19 16 16",
  "72 73 70 68 65 61",
  "27 30 28 25 24 23 18",
  "48 50 49 47 50 48 47",
  "93 96 99 96 99",
  "14 16 15 16 13 10 7 7",
  "77 79 77 74 77 76 72",
  "57 59 56 58 56 54 47",
  "74 77 75 75 72",
  "20 22 20 18 17 17 14 15",
  "57 59 58 56 56 56",
  "41 43 42 40 38 38 35 31",
  "96 98 97 96 96 95 88",
  "56 58 57 53 50 48",
  "76 77 75 71 73",
  "82 84 80 77 74 73 70 70",
  "66 67 66 62 58",
  "46 47 43 41 39 33",
  "58 60 59 57 51 50 48",
  "64 65 62 55 52 53",
  "37 38 37 31 29 26 26",
  "43 44 39 38 35 31",
  "79 80 75 73 67",
  "37 37 34 33 30 28",
  "65 65 64 63 62 63",
  "60 60 59 58 56 56",
  "83 83 81 78 76 74 73 69",
  "23 23 20 18 13",
  "10 10 13 12 10",
  "69 69 68 69 67 66 67",
  "13 13 11 14 13 10 8 8",
  "78 78 75 77 76 73 72 68",
  "79 79 76 73 75 69",
  "89 89 89 88 86 84 82 81",
  "62 62 59 59 56 57",
  "80 80 77 76 76 76",
  "99 99 97 94 94 90",
  "62 62 59 56 53 53 48",
  "69 69 65 62 61 59 58",
  "13 13 12 11 8 4 5",
  "69 69 67 66 64 60 60",
  "56 56 53 52 48 44",
  "61 61 57 56 50",
  "61 61 60 59 53 51",
  "82 82 81 78 77 70 67 69",
  "90 90 88 85 79 77 77",
  "91 91 86 85 81",
  "25 25 23 20 17 12 11 5",
  "15 11 8 7 6 4",
  "34 30 29 28 26 25 22 24",
  "29 25 24 21 20 17 14 14",
  "50 46 44 41 40 36",
  "52 48 46 43 38",
  "73 69 71 68 66 63 62 60",
  "45 41 38 37 36 37 39",
  "85 81 79 78 79 79",
  "38 34 33 31 32 30 29 25",
  "80 76 74 75 68",
  "48 44 44 41 39 37 34",
  "46 42 39 36 35 35 32 33",
  "61 57 54 52 52 50 50",
  "63 59 57 56 56 52",
  "31 27 26 26 23 18",
  "42 38 34 33 32 29 27",
  "37 33 31 28 25 21 18 21",
  "57 53 50 46 46",
  "24 20 18 14 13 9",
  "73 69 68 65 64 61 57 50",
  "74 70 68 65 64 63 57 54",
  "89 85 84 78 80",
  "54 50 44 42 42",
  "31 27 22 19 15",
  "35 31 24 22 20 17 10",
  "65 58 57 55 54",
  "36 31 30 29 27 26 29",
  "30 23 20 19 19",
  "27 22 21 20 16",
  "74 68 67 64 62 55",
  "12 5 3 6 5",
  "54 48 51 49 46 44 45",
  "24 18 19 16 13 13",
  "28 22 19 22 19 18 14",
  "61 55 52 53 48",
  "88 82 80 80 78 75 73",
  "93 87 86 86 87",
  "18 11 8 8 5 3 2 2",
  "87 80 79 78 77 77 73",
  "47 41 40 40 33",
  "96 89 85 83 82",
  "56 51 50 46 44 47",
  "82 76 74 70 68 66 65 65",
  "41 36 33 31 27 25 23 19",
  "77 72 70 66 63 61 56",
  "48 41 40 34 31 30",
  "80 74 69 66 65 63 65",
  "55 50 44 43 40 38 35 35",
  "81 74 71 69 66 59 56 52",
  "56 51 48 43 36",
  "55 58 59 62 65 68 69 68",
  "79 82 85 88 90 92 94 94",
  "62 63 65 67 70 74",
  "51 53 56 59 60 67",
  "42 43 41 43 44",
  "46 48 49 52 53 54 52 49",
  "79 81 82 79 82 82",
  "16 18 19 18 21 24 28",
  "32 33 35 37 40 39 40 45",
  "18 19 19 21 22 24 27",
  "5 8 10 10 12 9",
  "53 55 58 58 60 61 62 62",
  "83 85 88 90 91 94 94 98",
  "31 32 34 35 35 38 41 48",
  "11 12 16 17 20 21 23 24",
  "22 25 29 30 31 32 29",
  "57 59 61 65 65",
  "39 41 42 46 49 52 56",
  "37 38 42 44 50",
  "38 40 43 46 53 54",
  "59 62 68 70 68",
  "29 30 37 38 39 39",
  "74 75 76 83 87",
  "36 39 46 48 51 57",
  "38 35 36 39 41 43",
  "9 7 10 12 9",
  "49 46 48 49 50 50",
  "53 52 55 56 57 58 62",
  "50 48 51 54 61",
  "88 87 88 90 92 89 91",
  "82 80 83 85 84 81",
  "77 75 74 75 77 77",
  "18 15 17 19 17 20 24",
  "70 69 71 69 71 73 78",
  "10 7 9 10 12 12 14 16",
  "74 73 76 78 78 80 78",
  "44 42 42 44 44",
  "80 79 79 80 84",
  "62 59 61 62 64 65 65 70",
  "5 4 8 9 12 13",
  "79 78 82 83 81",
  "72 70 71 75 77 80 80",
  "11 9 12 14 18 22",
  "51 48 52 53 60",
  "65 64 67 72 75",
  "61 58 61 62 69 72 70",
  "41 38 41 42 48 49 51 51",
  "61 60 62 69 71 75",
  "10 9 15 17 18 24",
  "60 60 63 64 65",
  "66 66 68 70 71 70",
  "51 51 53 56 58 61 62 62",
  "43 43 46 49 53",
  "78 78 79 81 83 86 88 93",
  "69 69 70 73 76 78 75 77",
  "1 1 3 5 6 7 6 5",
  "41 41 40 43 43",
  "51 51 53 51 55",
  "57 57 55 56 61",
  "82 82 83 84 86 86 88",
  "89 89 92 92 91",
  "66 66 69 69 69",
  "69 69 71 74 75 75 77 81",
  "41 41 41 44 47 50 51 57",
  "12 12 15 19 20",
  "57 57 61 63 60",
  "53 53 56 59 63 63",
  "36 36 40 41 45",
  "82 82 86 88 89 90 95",
  "23 23 25 26 32 35",
  "77 77 78 79 85 82",
  "70 70 72 79 80 80",
  "32 32 34 35 40 44",
  "66 66 67 69 72 75 82 88",
  "79 83 85 87 89 90 91",
  "80 84 85 88 86",
  "90 94 97 99 99",
  "71 75 76 78 82",
  "14 18 20 23 25 28 35",
  "42 46 49 46 48 49 51",
  "72 76 77 74 77 80 81 78",
  "34 38 36 39 39",
  "71 75 74 76 78 82",
  "34 38 39 41 40 43 46 51",
  "48 52 52 53 56 57 60 62",
  "66 70 71 74 74 72",
  "53 57 57 59 60 63 63",
  "52 56 58 58 62",
  "5 9 12 12 15 16 19 25",
  "32 36 39 43 45",
  "18 22 25 26 28 32 33 32",
  "21 25 26 30 30",
  "47 51 54 58 60 64",
  "70 74 75 76 79 83 90",
  "36 40 46 49 52 55",
  "59 63 66 68 74 76 74",
  "53 57 58 59 61 62 69 69",
  "33 37 39 45 49",
  "35 39 46 49 51 52 55 61",
  "32 39 42 44 47 50",
  "69 74 77 80 83 81",
  "11 16 19 22 25 27 27",
  "58 63 66 68 70 74",
  "37 43 46 47 50 52 53 58",
  "84 89 88 91 92 95 97 99",
  "62 69 71 74 71 74 71",
  "4 10 8 9 11 14 14",
  "60 66 65 67 71",
  "39 46 45 46 47 52",
  "6 12 13 13 14 17",
  "36 42 42 45 44",
  "6 11 13 13 16 19 20 20",
  "34 41 41 44 45 49",
  "78 85 88 88 94",
  "72 78 81 85 86 89 91 94",
  "51 57 59 62 66 67 64",
  "51 57 59 61 65 67 67",
  "19 25 29 30 32 36",
  "44 50 51 55 57 64",
  "45 50 52 59 61",
  "12 17 20 23 26 29 34 31",
  "35 41 44 46 51 51",
  "28 34 35 38 40 42 48 52",
  "31 38 44 45 46 48 51 58",
  "74 72 70 69 66 63 62 64",
  "78 75 73 71 71",
  "66 65 63 62 60 56",
  "54 52 49 48 45 42 36",
  "65 62 61 59 57 58 56 54",
  "15 12 10 7 5 6 7",
  "26 25 23 21 23 20 20",
  "31 30 29 28 29 28 24",
  "35 33 31 34 31 28 25 19",
  "12 9 7 4 4 3 2",
  "99 97 95 92 91 91 88 90",
  "99 96 96 94 94",
  "45 42 42 39 36 35 31",
  "50 48 48 45 43 41 39 33",
  "88 85 81 78 77 76 74",
  "28 27 25 21 19 18 17 19",
  "30 28 27 25 23 19 19",
  "43 40 38 34 31 28 24",
  "24 21 20 19 18 17 13 8",
  "40 39 33 31 29 26 23",
  "26 25 23 16 15 13 11 14",
  "35 34 28 27 27",
  "17 16 10 7 3",
  "43 40 35 34 31 28 25 18",
  "48 50 47 46 44 41 39 38",
  "30 32 31 28 26 27",
  "55 56 54 52 51 48 48",
  "47 48 45 44 41 38 36 32",
  "35 36 33 30 23",
  "47 48 46 45 44 45 44 41",
  "20 23 22 21 18 15 16 19",
  "64 67 70 68 66 66",
  "47 50 53 52 51 50 48 44",
  "32 34 35 32 29 28 23",
  "44 47 45 45 44",
  "24 27 27 24 25",
  "86 88 85 85 85",
  "12 14 14 13 12 10 7 3",
  "72 75 72 72 71 68 63",
  "97 99 98 96 95 94 90 88",
  "41 42 39 38 35 33 29 32",
  "11 14 11 10 6 5 5",
  "76 78 75 72 69 65 61",
  "14 15 12 8 2",
  "54 57 54 52 47 46",
  "51 54 52 49 46 45 40 43",
  "15 18 15 8 8",
  "87 90 88 86 84 77 75 71",
  "77 80 79 73 72 71 66",
  "23 23 21 18 15 14",
  "81 81 78 76 79",
  "28 28 26 23 21 18 18",
  "35 35 33 31 28 27 25 21",
  "97 97 94 92 89 87 86 80",
  "3 3 5 4 2",
  "72 72 71 74 72 75",
  "44 44 43 40 41 39 37 37",
  "44 44 41 38 39 38 34",
  "90 90 89 92 90 84",
  "42 42 42 39 37 35 34",
  "61 61 61 59 57 56 57",
  "96 96 95 95 93 90 87 87",
  "91 91 90 88 86 86 84 80",
  "88 88 88 86 85 84 82 76",
  "99 99 95 94 92",
  "58 58 55 54 50 53",
  "17 17 13 11 11",
  "51 51 48 47 43 40 38 34",
  "92 92 88 85 82 77",
  "94 94 87 86 85",
  "15 15 13 11 10 4 7",
  "70 70 65 64 63 61 61",
  "55 55 49 46 42",
  "84 84 83 82 76 75 74 68",
  "55 51 48 45 42 40 37",
  "14 10 9 7 4 1 4",
  "33 29 26 23 22 21 21",
  "58 54 51 50 46",
  "17 13 10 8 3",
  "73 69 66 63 61 63 60 59",
  "35 31 29 30 31",
  "25 21 18 19 18 16 16",
  "72 68 71 70 69 65",
  "78 74 76 73 72 65",
  "86 82 80 79 76 76 75 74",
  "19 15 14 14 12 13",
  "86 82 80 80 78 75 75",
  "42 38 36 36 33 29",
  "47 43 43 41 38 35 33 26",
  "86 82 80 76 73 71 68",
  "62 58 56 52 51 48 47 50",
  "62 58 56 54 50 47 47",
  "20 16 15 14 10 9 5",
  "49 45 41 39 37 36 33 26",
  "51 47 46 45 39 36 34 32",
  "90 86 85 79 80",
  "25 21 19 17 10 8 6 6",
  "29 25 22 15 11",
  "62 58 57 54 49 44",
  "68 63 61 59 56 55",
  "47 41 39 38 37 35 37",
  "45 38 35 32 29 27 26 26",
  "84 77 74 71 68 64",
  "43 38 35 32 31 26",
  "25 19 17 16 17 16 13 10",
  "68 61 59 58 56 58 55 58",
  "54 47 46 43 46 44 43 43",
  "40 33 32 29 32 29 28 24",
  "77 70 69 72 71 69 66 61",
  "53 47 47 45 43 41 38",
  "80 75 72 71 71 69 72",
  "40 33 30 27 27 24 23 23",
  "22 16 13 13 9",
  "26 20 20 18 15 9",
  "55 48 45 41 38 36",
  "23 18 17 13 10 9 8 10",
  "25 19 15 12 12",
  "40 33 32 31 29 27 23 19",
  "93 88 84 81 78 73",
  "72 65 59 58 55 54 51 49",
  "64 58 52 49 52",
  "49 43 40 33 32 29 29",
  "63 57 56 54 47 43",
  "82 77 74 68 66 60",
  "58 55 53 46 46",
  "75 80 82 84 83 85 85",
  "23 19 21 20 17 14 14",
  "77 81 85 86 87 88 95",
  "47 41 39 37 36 37 31",
  "84 77 70 69 72",
  "82 81 79 76 76",
  "84 84 87 88 85 88 90",
  "76 74 72 70 67 70 68",
  "85 86 84 82 76 74 71 67",
  "87 87 84 83 82 78 74",
  "5 4 7 10 13 18 19",
  "7 5 7 4 5",
  "77 77 75 74 70",
  "73 69 66 63 60 59 56",
  "41 41 43 46 49 50 53 53",
  "46 50 52 54 54 56 60",
  "64 64 65 71 74",
  "52 55 52 50 48 45 42 44",
  "39 36 37 44 46 45",
  "59 55 54 52 48 45 41",
  "85 82 81 82 80 77 76 72",
  "61 56 54 51 48 45 44 40",
  "78 77 74 72 74 76",
  "8 13 17 18 21",
  "37 37 34 32 33",
  "46 44 45 48 49 55 55",
  "74 75 79 80 82 84 86 90",
  "86 82 82 80 79 76 74",
  "61 68 70 72 74 72",
  "9 9 13 15 22",
  "39 38 39 39 43",
  "2 9 7 10 14",
  "16 15 14 10 12",
  "33 37 44 46 48",
  "39 35 31 28 28",
  "61 61 63 64 64 65 65",
  "67 60 58 55 54 52 51 50",
  "69 70 68 67 68 65 62",
  "42 41 38 37 37 35 33 29",
  "64 66 70 73 76 79 80",
  "18 22 23 25 26 29 36",
  "37 36 29 28 26 25 26",
  "52 45 43 39 38 36 38",
  "51 56 57 63 65 67 69 75",
  "16 17 14 10 8 10",
  "20 14 14 12 9 6",
  "25 21 19 18 18 16 14 9",
  "40 40 37 35 34 33 31 24",
  "40 42 47 49 52 52",
  "72 73 72 68 66",
  "95 92 91 88 87 84 80",
  "76 75 74 72 70 68 64 60",
  "49 49 50 53 55 58 61 68",
  "50 56 59 62 66 68 69 69",
  "85 85 83 84 87 90 94",
  "21 22 18 17 11",
  "37 40 38 41 47",
  "18 19 16 13 10 10",
  "45 46 52 53 54 55 57 61",
  "2 5 5 8 11 9",
  "42 38 37 40 41",
  "28 25 23 25 22 17",
  "53 54 52 50 47 50 46",
  "18 22 25 26 27 27",
  "62 61 60 63 63",
  "30 30 31 38 41 41",
  "97 97 95 94 93 90 88",
  "32 33 31 25 24 21 20",
  "41 43 46 46 47 49",
  "25 24 27 29 31 32 34",
  "2 4 3 4 7 7",
  "26 23 21 19 17 17 16",
  "59 62 64 67 70 74 74",
  "62 62 60 57 50 49 47 48",
  "49 46 45 40 33",
  "44 38 36 37 37",
  "87 81 79 79 76 76",
  "7 14 17 19 19",
  "80 82 78 77 74 72 68",
  "5 9 11 13 19 21 21",
  "38 40 42 44 50 52 57",
  "74 78 79 81 83 82 84 89",
  "91 94 93 88 89",
  "24 28 31 33 36 39 44 43",
  "26 24 21 16 12",
  "86 82 81 75 73 71 69 69",
  "25 25 26 28 25",
  "80 79 83 86 87",
  "69 68 71 73 76 79 81 85",
  "18 16 19 22 22 24 31",
  "55 58 54 53 50 47 44 44",
  "37 31 30 24 23 23",
  "16 19 17 17 13",
  "58 56 55 54 54 51 51",
  "73 73 76 77 76 74",
  "21 24 21 18 13",
  "71 76 79 81 81 83 82",
  "37 39 38 38 35 34 32 26",
  "79 79 78 76 73 66 62",
  "46 44 45 47 48 53 54 58",
  "53 55 56 62 63 62",
  "20 23 20 18 18 18",
  "69 69 67 65 68 65 61",
  "50 50 51 51 56",
  "82 85 85 84 83 81 82",
  "71 67 70 69 66 65 62 55",
  "52 45 43 40 39 36 32 31",
  "86 89 91 93 93 96 96",
  "71 70 69 66 66 65 59",
  "13 13 10 9 8 8 4",
  "26 30 34 37 39",
  "38 41 39 36 33 30 30 29",
  "27 26 30 33 36 37 42",
  "34 36 34 32 31 29 30 33",
  "31 31 33 34 38 39 41 44",
  "98 92 92 90 87 88",
  "34 33 31 30 28 25 21 21",
  "77 71 73 71 69 65",
  "61 61 61 64 67 69 72",
  "64 68 70 71 72 75",
  "29 22 18 17 16 9",
  "43 46 49 52 56 55",
  "27 20 19 17 20",
  "51 47 44 43 40 36",
  "77 74 71 68 67 66 67",
  "46 52 55 57 61 62 63 61",
  "84 80 75 73 71 69 67 63",
  "63 63 66 63 63",
  "55 62 67 69 71 73 76 77",
  "82 81 84 86 88 91 90",
  "73 77 80 83 84 85 83 87",
  "47 50 47 40 37 35 32 25",
  "73 73 67 66 66",
  "90 90 93 92 90 91",
  "53 48 45 42 39 36 37 39",
  "51 47 41 40 39 36 30",
  "60 64 65 67 71 74 74",
  "25 27 29 29 33",
  "33 33 34 40 44",
  "68 73 74 75 79 83",
  "20 23 20 18 15",
  "12 15 17 19 21 21 23 29",
  "21 23 26 29 31 32 32",
  "86 86 84 86 86",
  "22 19 16 18 21 23 20",
  "30 31 28 26 22",
  "98 94 89 86 84",
  "6 11 13 10 16",
  "84 87 84 83 84 81 79 73",
  "39 35 34 34 33 32 30 31",
  "54 48 46 43 40 38 38",
  "9 13 11 13 15",
  "7 8 9 12 14 16 17 23",
  "27 27 30 32 33 36 38 42",
  "61 56 53 54 51",
  "21 18 20 21 24 30",
  "84 88 88 90 92 94 96 96",
  "66 66 64 61 60 60",
  "52 52 54 54 58",
  "68 64 63 59 58 57 52",
  "70 70 65 62 60",
  "55 48 47 45 39 32",
  "46 53 54 55 58 62",
  "64 64 64 63 60 53",
  "60 66 69 71 74 77 82 82",
  "6 11 12 10 7",
  "9 10 7 4 6 6",
  "11 9 12 16 15",
  "75 72 72 75 75",
  "55 55 52 49 45 43 40 39",
  "27 33 35 36 39 45 49",
  "61 61 54 53 47",
  "92 91 92 95 95 94",
  "35 35 37 39 43 43",
  "28 23 22 20 18 18 11",
  "38 42 45 48 47",
  "5 7 9 11 17 19 22",
  "48 44 42 41 40 41 39 35",
  "38 38 37 33 36",
  "69 65 65 63 61 60 56",
  "28 25 26 25 28 32",
  "25 23 26 30 34",
  "57 57 58 63 65 67 64",
  "42 43 45 47 44 45 46 50",
  "48 44 41 39 39",
  "71 65 59 56 54 52 51 47",
  "11 8 7 7 10",
  "59 56 55 57 54 51 50 50",
  "66 65 70 72 73 75 78 85",
  "34 37 39 42 43 45 46",
  "67 64 61 58 55 54 51 50",
  "1 3 5 7 9 10",
  "87 90 92 94 97",
  "48 50 51 54 55 58 59",
  "43 45 46 48 50",
  "35 32 30 28 27",
  "57 55 54 53 52",
  "52 54 55 57 60 63",
  "34 37 38 39 40 41 44",
  "75 78 81 84 86 88 90 93",
  "59 58 56 53 52 49 47",
  "23 26 27 28 29 30 32",
  "29 30 33 34 37 38 39",
  "50 48 46 44 43 42 39",
  "13 11 9 6 5",
  "70 72 74 76 79 81",
  "18 21 24 25 26",
  "21 18 15 12 10 9",
  "62 60 57 56 55",
  "13 11 8 7 6 5",
  "40 38 37 36 35",
  "47 50 52 55 57 59 60",
  "76 74 71 70 67 64 61",
  "39 40 43 46 49 52 55 56",
  "85 87 89 90 92",
  "35 33 32 30 27 25 24 23",
  "64 63 60 58 55 52 51",
  "82 80 78 76 73 71",
  "78 80 83 85 87 88 90 91",
  "40 39 38 35 34 33",
  "24 22 21 19 18 17 16",
  "14 17 18 21 23",
  "34 35 37 38 40 41 44 47",
  "97 96 93 92 90 88",
  "71 74 76 77 80 82",
  "54 57 59 60 63",
  "32 30 29 28 26",
  "41 38 36 33 32 30 28",
  "68 71 74 76 77 80",
  "73 71 69 67 64 62 59 56",
  "82 84 85 88 90",
  "38 39 40 42 44 47 49 52",
  "92 89 86 84 81 80",
  "61 62 64 67 69",
  "35 32 30 28 26 24 23 22",
  "14 16 18 19 22 25 28",
  "67 68 69 72 75 76 77 80",
  "29 28 27 25 22 19",
  "14 13 12 9 7 6",
  "22 24 27 28 30 33",
  "64 65 68 69 70 73 75 77",
  "66 65 63 60 59 56",
  "73 74 77 78 81",
  "12 13 16 17 20",
  "45 48 50 51 54 57 60 63",
  "45 43 41 38 35 33 30",
  "61 58 56 55 52 50",
  "61 59 56 53 50 49 46",
  "40 42 44 45 48 49 50 52",
  "75 77 80 81 83 85",
  "14 17 18 20 21 23 25",
  "13 12 9 8 7",
  "55 52 51 50 47 44 42",
  "64 63 60 59 56 54 53 52",
  "19 18 17 16 13 12 10 7",
  "99 97 94 91 89 86 83 80",
  "85 86 87 90 93 94 95 98",
  "16 19 20 23 25 26 27",
  "14 13 10 9 7",
  "83 82 79 77 76 74 71 69",
  "84 83 82 80 78 76",
  "20 21 24 26 29 30 31",
  "84 83 80 78 77 76 75",
  "32 34 35 36 38",
  "12 13 15 18 19 22 23 25",
  "56 53 52 49 46 43 42 40",
  "72 71 69 67 64 61 58 56",
  "33 36 38 39 40 43",
  "5 7 8 10 12 15 17 19",
  "40 42 43 46 49 52 53 55",
  "35 34 31 30 29 27 26 25",
  "32 33 35 36 37 40 41 43",
  "56 59 60 61 64 65 66",
  "55 58 60 62 64 65 68 71",
  "24 25 27 29 30 31 33 36",
  "97 95 92 89 88 85 83 81",
  "27 30 32 34 35 36",
  "17 18 21 22 25",
  "52 50 48 47 44 42 40 39",
  "5 7 10 11 13 16",
  "65 66 68 69 72 73 76 79",
  "72 71 68 67 64 61 59 57",
  "57 60 61 62 63 65 67",
  "86 85 84 82 80 79 76 74",
  "36 33 31 28 27 25 23",
  "21 22 23 26 29 31",
  "35 36 39 41 43 44 46",
  "73 70 67 64 63 61 58",
  "37 34 31 30 29 27",
  "69 72 75 78 79 81 82 83",
  "38 37 35 33 32 29 26 23",
  "60 59 58 55 54 53",
  "54 53 50 49 48 46 45",
  "86 87 90 93 94 96 98",
  "99 98 97 94 93 91",
  "47 46 44 43 40",
  "38 40 42 44 47 49",
  "86 88 90 93 96",
  "34 31 30 29 26",
  "43 46 47 50 51 54",
  "56 58 61 64 66 68 70 72",
  "62 65 67 69 72 75 78",
  "13 16 19 22 23 24 26",
  "78 80 81 82 85 87 89",
  "36 34 32 30 29 27 25 22",
  "36 35 34 31 30 28 25 24",
  "77 76 75 73 70 68 67",
  "20 17 15 12 10 9 6 4",
  "67 70 73 75 77",
  "52 55 57 58 60 63 66",
  "27 28 31 34 35",
  "17 18 21 23 26 29 30 33",
  "40 39 36 33 32 29 26",
  "11 8 7 5 2",
  "4 6 8 11 14 15",
  "94 92 89 87 84 83 80",
  "72 73 75 76 79 81 82 84",
  "26 29 32 34 37 38",
  "19 16 13 11 10 9 8",
  "79 82 85 87 89",
  "22 25 28 30 31 33 34 37",
  "89 86 83 82 81",
  "53 54 57 58 61 63",
  "88 85 84 83 82 79 77",
  "33 36 38 39 42 44",
  "37 39 40 41 44",
  "13 11 9 6 3",
  "25 26 29 30 32 34 36",
  "32 34 36 39 40 43 44 45",
  "26 24 21 19 18",
  "62 60 58 56 53 51 49 48",
  "37 35 33 30 27 25 22 21",
  "11 12 15 18 19 20",
  "87 85 82 80 78 76 75",
  "52 54 55 56 59",
  "83 82 80 77 76 74 73",
  "81 79 77 74 72 71 68 67",
  "80 81 83 84 87",
  "11 13 16 18 21 24 25",
  "42 40 39 37 35 34",
  "19 18 16 14 13 12 9",
  "89 87 85 84 81",
  "91 90 88 86 83",
  "85 82 80 77 76 74 71 70",
  "52 54 57 59 62",
  "2 4 7 8 10",
  "61 64 66 67 68 71 73",
  "29 30 33 36 38 39 41",
  "58 57 54 52 51 49",
  "83 84 85 87 89 91",
  "23 22 19 16 14 13 11",
  "41 42 44 45 48 51 52 53",
  "32 30 28 26 23 20 17",
  "72 73 75 76 77",
  "78 75 74 71 68 65",
  "87 88 89 91 93 95",
  "81 84 85 86 89 91",
  "26 25 24 22 19 18 16",
  "44 42 40 39 36 34",
  "63 65 67 69 72 73",
  "48 51 53 56 58 61 64",
  "95 93 90 87 84 82 79 77",
  "45 46 49 52 54 55",
  "36 33 30 28 27 24",
  "4 6 7 9 12 15",
  "20 19 17 14 13 10 9 7",
  "6 8 11 12 14",
  "52 53 55 57 58 61 64 67",
  "58 55 54 52 50 49",
  "89 88 85 83 82",
  "15 12 11 10 9",
  "30 29 27 25 22 20 19 16",
  "63 65 66 67 70",
  "62 60 57 55 54 51 49",
  "49 48 47 45 44 42 40",
  "91 88 86 83 80 77 75",
  "54 57 59 60 63 65",
  "89 92 94 95 97",
  "6 9 11 12 15 16",
  "91 89 87 86 84 83 82",
  "41 42 43 44 46 47 48 49",
  "40 37 34 33 30 29",
  "36 39 40 43 44 46",
  "53 55 58 59 60 63 66",
  "69 70 73 74 76 77",
  "59 60 61 63 64 66 69 70",
  "12 13 15 16 19 22 23",
  "94 93 92 90 89 88 87 86",
  "1 4 7 10 13 15 18",
  "44 42 40 38 35 34 31 28",
  "99 98 96 94 93",
  "7 9 10 13 16 19 20 22",
  "80 79 76 75 74",
  "79 81 83 84 87 88",
  "76 73 72 70 68 67 64",
  "81 84 85 86 87 89 92 93",
  "53 55 56 58 59 62 64 65",
  "75 74 72 71 70 68",
  "33 30 28 25 22 20 17",
  "36 33 30 27 24 22",
  "76 77 78 80 81 84 87",
  "72 69 68 66 64 62",
  "33 35 38 39 42",
  "53 56 59 60 62 63 65",
  "46 47 48 51 53",
  "16 17 19 20 23 25 28",
  "74 76 79 82 84",
  "29 26 23 20 18 16 14 12",
  "66 69 72 74 75 76",
  "32 34 36 38 41 44",
  "46 44 41 39 38",
  "59 60 61 62 64 67 68 71",
  "38 40 42 43 46",
  "45 43 42 39 38 35 32 30",
  "56 59 61 64 66 69 70 71",
  "2 5 7 10 11",
  "64 63 62 61 58 57",
  "46 44 41 39 38 37 35 32",
  "32 31 29 26 23 21",
  "97 96 94 91 89 88 85",
  "54 52 50 47 44 43 41 39",
  "20 18 15 14 12 10 9",
  "33 35 36 38 39 42 45 46",
  "15 17 19 21 24 25 26 27",
  "85 87 88 90 91 92",
  "61 58 57 54 51 48 46 45",
  "47 48 49 52 54",
  "29 27 24 23 20 19 18 16",
  "22 20 18 17 16 15 14 12",
  "65 66 67 69 71 74 76 77",
  "19 18 15 13 12 9 6",
  "25 22 19 17 15 14 12 9",
  "20 23 24 26 29",
  "35 36 39 41 43 45",
  "34 32 29 28 26",
  "15 18 19 20 21 22 23",
  "43 46 48 49 51",
  "84 82 81 79 77",
  "47 49 50 53 56 59",
  "22 19 18 17 14 12 11 8",
  "84 82 79 76 73 70 68",
  "51 48 46 43 41 39 37 36",
  "19 21 23 24 27 28 31",
  "60 58 55 54 51 49",
  "76 73 71 70 68 67",
  "18 15 13 11 10 7 5 3",
  "97 95 93 92 90",
  "46 47 49 52 55",
  "92 90 87 86 85 82",
  "80 81 83 84 86 88 90 92",
  "28 31 33 34 36 39",
  "44 47 50 52 55 57 60",
  "14 15 16 17 20",
  "99 97 96 93 90",
  "48 51 53 54 56 57 58 61",
  "86 87 90 91 94 97",
  "81 80 79 78 77 74",
  "62 63 66 68 69 72 75",
  "95 94 91 88 85 82",
  "40 39 37 36 33 31",
  "36 38 40 42 45",
  "71 69 66 63 60 59",
  "69 70 71 73 76",
  "71 74 75 76 79 80 82 85",
  "77 78 81 82 83 85 87",
  "14 11 9 6 5 2",
  "38 36 35 34 31 30",
  "71 69 66 63 60",
  "77 78 79 81 82",
  "23 22 19 16 14 11 10 8",
  "75 77 80 83 84 87 88",
  "64 65 67 69 71 74 75 77",
  "88 87 86 84 81 80 77 75",
  "48 47 44 41 40",
  "75 77 80 83 86 87 88 90",
  "34 35 38 39 41",
  "81 80 78 75 72 69",
  "62 61 60 58 56",
  "25 27 28 30 31",
  "50 47 45 43 42 39",
  "10 12 15 16 17 19 22 24",
  "50 51 52 54 56 57 59",
  "44 41 40 37 36 34 32",
  "76 73 70 67 66 64 63",
  "91 88 85 82 79 76 75 74",
  "81 82 84 86 88 91 92",
  "17 14 12 9 8 7 5",
  "24 23 21 19 16 13 10",
  "78 76 74 71 69 66",
  "61 63 66 68 71 72 73 75",
  "93 91 89 88 87 84",
  "23 25 27 30 31 33 36 38",
  "35 36 37 40 41 43 44 47",
  "81 84 85 86 89 91 93 95",
  "89 92 93 94 95",
  "37 38 39 42 45 46 48 51",
  "38 35 33 31 28 25 24 23",
  "38 36 34 32 31 30 28",
  "2 4 5 7 10 13 15 17",
  "23 26 29 32 34 35",
  "18 17 15 13 12 10 7",
  "51 48 45 42 40",
  "15 16 19 20 23 26 27",
}
