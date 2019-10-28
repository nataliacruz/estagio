#!usr/bin/perl

use strict;
use warnings;
use POSIX;

my $count;
my $bases;
open(my $filenameIn,"<",$ARGV[0]) or die "Can't open $ARGV[0]";
open(my $filenameOut, ">>", $ARGV[1]) or die "Can't open $ARGV[1]";
while(my $line = <$filenameIn>){
	$bases="";
	$count=0;
	chomp($line);
	my $sub = substr($line, 0, 1);
	if(($sub eq "#")==1){
                print $filenameOut "$line\n";
        }else{
                my @l = split("\t",$line);
		my @spl = split(",",$l[4]);
                my $spli = @spl; 
		if($spli > 1){
			foreach my $names (@spl){
				if(($names eq "<X>")==0){
					if($count==0){
						$bases = $names;
					}else{
						$count++;
						$bases = $bases .",". $names;
					}
				}
			}
			$l[4]=$bases;
			foreach my $col (@l){
				print $filenameOut "$col\t";
			}
                	print $filenameOut "\n";	
		}
	}	
}
close $filenameIn;
close $filenameOut;


####
