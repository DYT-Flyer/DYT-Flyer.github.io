package main

var simpleFiles [][]string = [][]string{
	{"1-1.jpg", "1-2.jpg", "1-3.png", "1-4.jpeg", "1-5.png", "1-6.jpeg", "1-7.jpg", "1-8.png"},
	{"2-1.jpg", "2-2.jpeg", "2-3.jpeg", "2-4.png", "2-5.png", "2-6.jpeg", "2-7.jpg", "2-8.png"},
	{"3-1.jpg", "3-2.jpeg", "3-3.png", "3-4.jpeg", "3-5.png", "3-6.jpeg", "3-7.jpg", "3-8.png"},
	{"4-1.jpg", "4-2.jpeg", "4-3.png", "4-4.jpeg", "4-5.png", "4-6.jpeg", "4-7.jpg", "4-8.png"},
	{"5-1.jpg", "5-2.jpeg", "5-3.png", "5-4.jpeg", "5-5.png", "5-6.jpeg", "5-7.jpg", "5-8.png"},
	{"6-1.jpg", "6-2.jpeg", "6-3.png", "6-4.jpeg", "6-5.png", "6-6.jpeg", "6-7.jpg", "6-8.png"},
	{"7-1.jpg", "7-2.jpeg", "7-3.png", "7-4.jpeg", "7-5.png", "7-6.jpeg", "7-7.jpg", "7-8.png"},
	{"8-1.jpg", "8-2.jpeg", "8-3.png", "8-4.jpeg", "8-5.png", "8-6.jpeg", "8-7.jpg", "8-8.png"},
	{"9-1.jpg", "9-2.jpeg", "9-3.png", "9-4.jpeg", "9-5.png", "9-6.jpeg", "9-7.jpg", "9-8.png"},
	{"10-1.jpg", "10-2.jpeg", "10-3.png", "10-4.jpeg", "10-5.png", "10-6.jpeg", "10-7.jpg", "10-8.png"},
	{"11-1.jpg", "11-2.jpg", "11-3.png", "11-4.jpg", "11-5.jpg", "11-6.jpg", "11-7.jpg", "11-8.png"},
	{"12-1.jpg", "12-2.jpg", "12-3.png", "12-4.jpg", "12-5.jpg", "12-6.jpg", "12-7.jpg", "12-8.png"},
	{"13-1.jpg", "13-2.jpg", "13-3.png", "13-4.jpg", "13-5.jpg", "13-6.jpg", "13-7.jpg", "13-8.png"},
	{"14-1.jpg", "14-2.jpg", "14-3.png", "14-4.jpg", "14-5.jpg", "14-6.jpg", "14-7.jpg", "14-8.png"},
	{"15-1.jpg", "15-2.jpeg", "15-3.png", "15-4.jpg", "15-5.png", "15-6.jpg", "15-7.jpg", "15-8.png"},
	{"16-1.jpg", "16-2.jpeg", "16-3.png", "16-4.jpg", "16-5.png", "16-6.jpeg", "16-7.jpg", "16-8.png"},
	{"17-1.jpg", "17-2.jpeg", "17-3.png", "17-4.jpg", "17-5.png", "17-6.jpeg", "17-7.jpg", "17-8.png"},
	{"18-1.jpg", "18-2.jpeg", "18-3.png", "18-4.jpg", "18-5.png", "18-6.jpeg", "18-7.jpg", "18-8.png"},
	{"19-1.jpg", "19-2.jpeg", "19-3.png", "19-4.jpg", "19-5.png", "19-6.jpeg", "19-7.jpg", "19-8.png"},
	{"20-1.jpg", "20-2.jpeg", "20-3.png", "20-4.jpg", "20-5.png", "20-6.jpeg", "20-7.jpg", "20-8.png"},
	{"21-1.jpg", "21-2.jpeg", "21-3.png", "21-4.jpg", "21-5.png", "21-6.jpeg", "21-7.jpg", "21-8.png"},
	{"22-1.jpg", "22-2.jpeg", "22-3.png", "22-4.jpg", "22-5.png", "22-6.jpeg", "22-7.jpg", "22-8.png"},
	{"23-1.jpg", "23-2.jpeg", "23-3.png", "23-4.jpg", "23-5.png", "23-6.jpeg", "23-7.jpg", "23-8.png"},
	{"24-1.jpg", "24-2.jpeg", "24-3.png", "24-4.jpg", "24-5.png", "24-6.jpeg", "24-7.jpg", "24-8.png"},
	{"25-1.jpg", "25-2.jpeg", "25-3.png", "25-4.jpg", "25-5.png", "25-6.jpeg", "25-7.jpg", "25-8.png"},
	{"26-1.jpg", "26-2.jpeg", "26-3.png", "26-4.jpeg", "26-5.png", "26-6.jpeg", "26-7.jpg", "26-8.png"},
	{"27-1.jpg", "27-2.jpeg", "27-3.png", "27-4.jpeg", "27-5.png", "27-6.jpeg", "27-7.jpg", "27-8.png"},
	{"28-1.jpg", "28-2.jpeg", "28-3.png", "28-4.jpeg", "28-5.png", "28-6.jpeg", "28-7.jpg", "28-8.png"},
	{"29-1.jpg", "29-2.jpeg", "29-3.png", "29-4.jpeg", "29-5.png", "29-6.jpeg", "29-7.jpg", "29-8.png"},
	{"30-1.jpg", "30-2.jpeg", "30-3.png", "30-4.jpeg", "30-5.png", "30-6.jpeg", "30-7.jpeg", "30-8.png"},
	{"31-1.jpg", "31-2.jpeg", "31-3.png", "31-4.jpeg", "31-5.png", "31-6.jpeg", "31-7.jpg", "31-8.png"},
	{"32-1.jpg", "32-2.jpeg", "32-3.png", "32-4.jpeg", "32-5.png", "32-6.jpeg", "32-7.jpg", "32-8.png"},
	{"33-1.jpg", "33-2.jpeg", "33-3.png", "33-4.jpeg", "33-5.png", "33-6.jpeg", "33-7.jpg", "33-8.png"},
	{"34-1.jpg", "34-2.jpeg", "34-3.png", "34-4.jpeg", "34-5.png", "34-6.jpeg", "34-7.jpg", "34-8.png"},
	{"35-1.jpg", "35-2.jpeg", "35-3.png", "35-4.jpeg", "35-5.png", "35-6.jpeg", "35-7.jpg", "35-8.png"},
	{"36-1.jpg", "36-2.jpeg", "36-3.png", "36-4.jpeg", "36-5.png", "36-6.jpeg", "36-7.jpg", "36-8.png"},
	{"37-1.jpg", "37-2.jpeg", "37-3.png", "37-4.jpeg", "37-5.png", "37-6.jpeg", "37-7.jpg", "37-8.png"},
	{"38-1.jpg", "38-2.jpeg", "38-3.png", "38-4.jpeg", "38-5.png", "38-6.jpeg", "38-7.jpg", "38-8.png"},
	{"39-1.jpg", "39-2.jpeg", "39-3.png", "39-4.jpeg", "39-5.png", "39-6.jpeg", "39-7.jpg", "39-8.png"},
	{"40-1.jpg", "40-2.jpeg", "40-3.png", "40-4.jpeg", "40-5.png", "40-6.jpeg", "40-7.jpg", "40-8.png"},
	{"41-1.jpg", "41-2.jpeg", "41-3.png", "41-4.jpeg", "41-5.png", "41-6.jpeg", "41-7.jpg", "41-8.png"},
	{"42-1.jpg", "42-2.jpeg", "42-3.png", "42-4.jpeg", "42-5.png", "42-6.jpeg", "42-7.jpg", "42-8.png"},
	{"43-1.jpg", "43-2.jpeg", "43-3.png", "43-4.jpeg", "43-5.png", "43-6.jpeg", "43-7.jpg", "43-8.png"},
	{"44-1.jpg", "44-2.jpeg", "44-3.png", "44-4.jpeg", "44-5.png", "44-6.jpeg", "44-7.jpg", "44-8.png"},
	{"45-1.jpg", "45-2.jpeg", "45-3.png", "45-4.jpeg", "45-5.png", "45-6.jpeg", "45-7.jpg", "45-8.png"},
	{"46-1.jpg", "46-2.jpeg", "46-3.png", "46-4.jpeg", "46-5.png", "46-6.jpeg", "46-7.jpg", "46-8.png"},
	{"47-1.jpg", "47-2.jpeg", "47-3.png", "47-4.jpeg", "47-5.png", "47-6.jpeg", "47-7.jpg", "47-8.png"},
	{"48-1.jpg", "48-2.jpeg", "48-3.png", "48-4.jpeg", "48-5.png", "48-6.jpeg", "48-7.jpg", "48-8.png"},
	{"49-1.jpg", "49-2.jpeg", "49-3.png", "49-4.jpeg", "49-5.png", "49-6.jpeg", "49-7.jpg", "49-8.png"},
	{"50-1.jpg", "50-2.jpeg", "50-3.png", "50-4.jpeg", "50-5.png", "50-6.jpeg", "50-7.jpg", "50-8.png"},
}

var complexFiles [][]string = [][]string{
	{"1-1.jpg", "1-2.jpeg", "1-3.png", "1-4.jpeg", "1-5.png", "1-6.jpeg", "1-7.jpg", "1-8.png"},
	{"2-1.jpg", "2-2.jpg", "2-3.jpeg", "2-4.jpeg", "2-5.png", "2-6.jpeg", "2-7.jpg", "2-8.png"},
	{"3-1.jpg", "3-2.jpeg", "3-3.png", "3-4.jpeg", "3-5.png", "3-6.jpeg", "3-7.jpg", "3-8.png"},
	{"4-1.jpg", "4-2.jpeg", "4-3.png", "4-4.jpeg", "4-5.png", "4-6.jpeg", "4-7.jpg", "4-8.png"},
	{"5-1.jpg", "5-2.jpeg", "5-3.png", "5-4.jpeg", "5-5.png", "5-6.jpeg", "5-7.jpg", "5-8.png"},
	{"6-1.jpg", "6-2.jpeg", "6-3.png", "6-4.jpeg", "6-5.png", "6-6.jpeg", "6-7.jpg", "6-8.png"},
	{"7-1.jpg", "7-2.jpeg", "7-3.png", "7-4.jpeg", "7-5.png", "7-6.jpeg", "7-7.jpg", "7-8.png"},
	{"8-1.jpg", "8-2.jpeg", "8-3.png", "8-4.jpeg", "8-5.png", "8-6.jpeg", "8-7.jpg", "8-8.png"},
	{"9-1.jpg", "9-2.jpeg", "9-3.png", "9-4.jpeg", "9-5.png", "9-6.jpeg", "9-7.jpg", "9-8.png"},
	{"10-1.jpg", "10-2.jpeg", "10-3.jpg", "10-4.jpeg", "10-5.png", "10-6.jpeg", "10-7.jpg", "10-8.png"},
	{"11-1.jpg", "11-2.jpeg", "11-3.jpg", "11-4.jpeg", "11-5.png", "11-6.jpeg", "11-7.jpg", "11-8.png"},
	{"12-1.jpg", "12-2.jpeg", "12-3.jpg", "12-4.jpeg", "12-5.png", "12-6.jpeg", "12-7.jpg", "12-8.png"},
	{"13-1.jpg", "13-2.jpeg", "13-3.jpg", "13-4.jpeg", "13-5.png", "13-6.jpeg", "13-7.jpg", "13-8.png"},
	{"14-1.jpg", "14-2.jpeg", "14-3.jpg", "14-4.jpeg", "14-5.png", "14-6.jpeg", "14-7.jpg", "14-8.png"},
	{"15-1.jpg", "15-2.jpeg", "15-3.jpg", "15-4.jpeg", "15-5.png", "15-6.jpeg", "15-7.jpg", "15-8.png"},
	{"16-1.jpg", "16-2.jpeg", "16-3.jpg", "16-4.jpeg", "16-5.png", "16-6.jpeg", "16-7.jpg", "16-8.png"},
	{"17-1.jpg", "17-2.jpeg", "17-3.jpg", "17-4.jpeg", "17-5.png", "17-6.jpeg", "17-7.jpg", "17-8.png"},
	{"18-1.jpg", "18-2.jpeg", "18-3.jpg", "18-4.jpeg", "18-5.png", "18-6.jpeg", "18-7.jpg", "18-8.png"},
	{"19-1.jpg", "19-2.jpeg", "19-3.jpg", "19-4.jpeg", "19-5.png", "19-6.jpeg", "19-7.jpg", "19-8.png"},
	{"20-1.jpg", "20-2.jpeg", "20-3.jpg", "20-4.jpeg", "20-5.png", "20-6.jpeg", "20-7.jpg", "20-8.png"},
	{"21-1.jpg", "21-2.jpeg", "21-3.jpg", "21-4.jpeg", "21-5.png", "21-6.jpeg", "21-7.jpg", "21-8.png"},
	{"22-1.jpg", "22-2.jpeg", "22-3.jpg", "22-4.jpeg", "22-5.png", "22-6.jpeg", "22-7.jpg", "22-8.png"},
	{"23-1.jpg", "23-2.jpeg", "23-3.jpg", "23-4.jpeg", "23-5.png", "23-6.jpeg", "23-7.jpg", "23-8.png"},
	{"24-1.jpg", "24-2.jpeg", "24-3.jpg", "24-4.jpeg", "24-5.png", "24-6.jpeg", "24-7.jpg", "24-8.png"},
	{"25-1.jpg", "25-2.jpg", "25-3.jpg", "25-4.jpg", "25-5.jpg", "25-6.jpg", "25-7.jpg", "25-8.png"},
	{"26-1.jpg", "26-2.jpg", "26-3.jpg", "26-4.jpg", "26-5.jpg", "26-6.jpg", "26-7.jpg", "26-8.png"},
	{"27-1.jpg", "27-2.jpeg", "27-3.jpg", "27-4.jpg", "27-5.jpg", "27-6.jpg", "27-7.jpg", "27-8.png"},
	{"28-1.jpg", "28-2.jpeg", "28-3.jpg", "28-4.jpg", "28-5.jpg", "28-6.jpg", "28-7.jpg", "28-8.png"},
	{"29-1.jpg", "29-2.jpeg", "29-3.jpg", "29-4.jpg", "29-5.jpg", "29-6.jpg", "29-7.jpg", "29-8.png"},
	{"30-1.jpg", "30-2.jpeg", "30-3.jpg", "30-4.jpg", "30-5.jpg", "30-6.jpg", "30-7.jpg", "30-8.png"},
	{"31-1.jpg", "31-2.jpg", "31-3.jpg", "31-4.jpg", "31-5.jpg", "31-6.jpg", "31-7.jpg", "31-8.png"},
	{"32-1.jpg", "32-2.jpg", "32-3.jpg", "32-4.jpg", "32-5.jpg", "32-6.jpg", "32-7.jpg", "32-8.png"},
	{"33-1.jpg", "33-2.jpg", "33-3.jpg", "33-4.jpg", "33-5.jpg", "33-6.jpg", "33-7.jpg", "33-8.png"},
	{"34-1.jpg", "34-2.jpg", "34-3.jpg", "34-4.jpg", "34-5.jpg", "34-6.jpg", "34-7.jpg", "34-8.png"},
	{"35-1.jpg", "35-2.jpg", "35-3.jpg", "35-4.jpg", "35-5.jpg", "35-6.jpg", "35-7.jpg", "35-8.png"},
	{"36-1.jpg", "36-2.jpg", "36-3.jpg", "36-4.jpg", "36-5.jpg", "36-6.jpg", "36-7.jpg", "36-8.png"},
	{"37-1.jpg", "37-2.jpg", "37-3.jpg", "37-4.jpg", "37-5.jpg", "37-6.jpg", "37-7.jpg", "37-8.png"},
	{"38-1.jpg", "38-2.jpg", "38-3.jpg", "38-4.jpg", "38-5.jpg", "38-6.jpg", "38-7.jpg", "38-8.png"},
	{"39-1.jpg", "39-2.jpg", "39-3.jpg", "39-4.jpg", "39-5.jpg", "39-6.jpg", "39-7.jpg", "39-8.png"},
	{"40-1.jpg", "40-2.jpg", "40-3.jpg", "40-4.jpg", "40-5.jpg", "40-6.jpg", "40-7.jpg", "40-8.png"},
	{"41-1.jpg", "41-2.jpg", "41-3.jpg", "41-4.jpg", "41-5.jpg", "41-6.jpg", "41-7.jpg", "41-8.png"},
	{"42-1.jpg", "42-2.jpg", "42-3.jpg", "42-4.jpg", "42-5.jpg", "42-6.jpg", "42-7.jpg", "42-8.png"},
	{"43-1.jpg", "43-2.jpg", "43-3.jpg", "43-4.jpg", "43-5.jpg", "43-6.jpg", "43-7.jpg", "43-8.png"},
	{"44-1.jpg", "44-2.jpg", "44-3.jpg", "44-4.jpg", "44-5.jpg", "44-6.jpg", "44-7.jpg", "44-8.png"},
	{"45-1.jpg", "45-2.jpg", "45-3.jpg", "45-4.jpg", "45-5.jpg", "45-6.jpg", "45-7.jpg", "45-8.png"},
	{"46-1.jpg", "46-2.jpg", "46-3.jpg", "46-4.jpg", "46-5.jpg", "46-6.jpg", "46-7.jpg", "46-8.png"},
	{"47-1.jpg", "47-2.jpg", "47-3.jpg", "47-4.jpg", "47-5.jpg", "47-6.jpg", "47-7.jpg", "47-8.png"},
	{"48-1.jpg", "48-2.jpg", "48-3.jpg", "48-4.jpg", "48-5.jpg", "48-6.jpg", "48-7.jpg", "48-8.png"},
	{"49-1.jpg", "49-2.jpg", "49-3.jpg", "49-4.jpg", "49-5.jpg", "49-6.jpg", "49-7.jpg", "49-8.png"},
	{"50-1.jpg", "50-2.jpg", "50-3.jpg", "50-4.jpg", "50-5.jpg", "50-6.jpg", "50-7.jpg", "50-8.png"},
}
