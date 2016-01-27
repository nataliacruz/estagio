from sys import argv
script, filenameIn, filenameOut = argv
bases = 0
q30 = 0
q20 = 0
medio = 0
i = count = 0
with open(filenameIn) as f:
	for line in f:
		line = line.strip("\n")
		if(i==1):
			maior = len(line)
			menor = len(line)	
		if(count==1):
			if(len(line)>maior):
				maior = len(line)
			if(len(line)<menor):
				menor = len(line)
			bases = bases+len(line)
		if(count==3):
			for letter in line:
				if(ord(letter)>30):
					q30 = q30+1
				if(ord(letter)>20):
					q20 = q20+1
			count=-1
		i = i+1
		count = count+1
reads = i/4
medio = bases/reads
file = open(filenameOut, 'a+')
file.write("%f\t %f\t %d\t %d\t %f\t %d\t %d\n" % (reads, bases, q30, q20, medio, maior, menor))
file.close()
