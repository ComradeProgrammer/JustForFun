import java.util.ArrayList;
import java.util.Scanner;

public class Item {
    private String word;
    private ArrayList<String> partOfSpeech;
    private ArrayList<String> meaning;
    private ArrayList<Integer>level;

    Item(String word){
        this.word=word;
    }

    public void addChineseMeaning(String partOfSpeech,String meaning,int level){
        this.partOfSpeech.add(partOfSpeech);
        this.meaning.add(meaning);
        this.level.add(level);
    }

    public void foreignToChineseExam(){
        System.out.println("你是否知道 "+word+" 的含义是什么,按回车继续");
        Scanner stdin=new Scanner(System.in);
        String input=stdin.nextLine();
        System.out.println(word+"记录在数据库的含义如下：");
        for(int i=0;i<meaning.size();i++){
            System.out.println(partOfSpeech.get(i)+" |"+meaning.get(i)+" | 目前等级"+level.get(i));
            System.out.println("请为这一条目重新打分");
            int newLevel=3;
            try {
                newLevel = Integer.parseInt(stdin.nextLine());
            }catch(Exception e){
                System.out.println("invalid input");
            }
            if(newLevel!=level.get(i)){
                modifyLevel(word,partOfSpeech.get(i),meaning.get(i),newLevel);
            }
        }
    }

    public void modifyLevel(String word,String partOfSpeech,String meaning,int newLevel){

    }

}
