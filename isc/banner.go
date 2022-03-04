package isc

import (
	"fmt"
	"io/ioutil"
)

var Banner = DefaultBanner

func PrintBanner() {
	fmt.Printf("%s\n", Banner)
}

func LoadBanner(filePath string) {
	if b, err := ioutil.ReadFile(filePath); err == nil {
		str := string(b)
		fmt.Printf("%s\n", str)
	} else {
		fmt.Printf("%s\n", Banner)
	}
}

var DefaultBanner = `
______   ____     __    __  ____     ____     _____    ____     ____      
/\__  _\ /\  _` + "`" + `\  /\ \  /\ \/\  _` + "`" + `\  /\  _` + "`" + `\  /\  __` + "`" + `\ /\  _` + "`" + `\  /\  _` + "`" + `\    
\/_/\ \/ \ \,\L\_\\ ` + "`" + `\` + "`" + `\\/'/\ \,\L\_\\ \ \/\_\\ \ \/\ \\ \ \L\ \\ \ \L\_\  
   \ \ \  \/_\__ \ ` + "`" + `\ ` + "`" + `\ /'  \/_\__ \ \ \ \/_/_\ \ \ \ \\ \ ,  / \ \  _\L  
    \_\ \__ /\ \L\ \ ` + "`" + `\ \ \    /\ \L\ \\ \ \L\ \\ \ \_\ \\ \ \\ \ \ \ \L\ \
    /\_____\\ ` + "`" + `\____\  \ \_\   \ ` + "`" + `\____\\ \____/ \ \_____\\ \_\ \_\\ \____/
    \/_____/ \/_____/   \/_/    \/_____/ \/___/   \/_____/ \/_/\/ / \/___/ 
`

var BannerMiku = `
               #########
              ############
              #############
             ##  ###########
            ###  ###### #####
            ### #######   ####
           ###  ########## ####
          ####  ########### ####
        #####   ###########  #####
       ######   ### ########   #####
       #####   ###   ########   ######
      ######   ###  ###########   ######
     ######   #### ##############  ######
    #######  ##################### #######
    #######  ##############################
   #######  ###### ################# #######
   #######  ###### ###### #########   ######
   #######    ##  ######   ######     ######
   #######        ######    #####     #####
    ######        #####     #####     ####
     #####        ####      #####     ###
      #####      ;###        ###      #
        ##       ####        ####

 :: iSysCore Service (GOLANG) ::

`

var BannerBinDwenDwen = `
        .:          :j        
         i.,,......:Ei:       
        KWG        .,#        
        Wi .:t.  i:, ;,       
       .. ..jj;;;;fKL .       
         iD:,     .iLi    .   
        ,i.         .Df .LWLi 
      . .: ###    ###L:::KEW  
      ;,; W#f#W  E:###fiLWKW  
      .::##t#iL ; K###DttW# : 
     ..; W###W  # ,,##,if#W.  
     , :.#### K .j ###tj,#    
     , i. ##    ..  K#;:;.t   
    ...,:.      ..  .:f; ;    
   .#L..,f..   ....:,;,;      
   K#W..t:;.:...::,,:Li:      
   W#K:.,t;it;j;,iLjfj;,      
  EW#D;... :iii,..ji,f,       
  D##:......,:,;;,:,:,,       
  :EL  :...        .:,.    :: Bing Dwen Dwen ::   
   .   .:..    i;  .::     :: iSysCore Service (GOLANG) ::    
       .::....   ..:,         
        ::::..:;;;:,.         
         f::::::j,,;L         
        ,LLL:,,,;ifG.         
        fDKEL;.:GWWW.         
        tWWWK   #WW#i         
        ,j,j.   t#Kjt
`
