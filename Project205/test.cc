#include<iostream>
#include<unistd.h>
#include<string>
#include<cstdlib>
#include <ctime> 
using namespace std;
int main(int argc, char* argv[]){
	string tmp=string(argv[1]);
	unsigned long seed = time(0);
    srand(seed);
	int t=rand()%10;
	cout<<"start "<<tmp<<" sleep "<<t<<endl;
	sleep(t);
	cout<<"end "<<tmp<<endl;
}