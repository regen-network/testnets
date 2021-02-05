1、Fork the repository(regen-network/testnets) to your VPS and make sure you are in the root directory  
```https://github.com/<your-github-username>/testnets``` 

2、Move to the folder where gentx is located   
```cd ~/.regen/config/gentx```

3、Move gentx files to the appropriate folder   
```cp <your-gentx-file-name> ~/testnets/aplikigo-1/gentxs/```

4、Verify if the file is in the folder   
```cd ~/testnets/aplikigo-1/gentxs/ ```    
```ls```   
***(Make sure you can find your gentx file!)***   

5、Save the change and submit the change      
```git add.```         
***(Don't forget the dot ".")***   

6、Add a note to the change     
```git commit -m "gentx submit"```        
***（Change anything inside the inverted commas as you want)***    

7、Post your push operation     
```git push```

8、Enter your github password and you're done!Great job!    
