package creationalPattern.factoryPattern;

public class GetPlanFactory {
    public static Plan getPlan(String planType) {
        Plan plan = null;
        switch(planType) {
            case "DOMESTIC" : {
                plan = new DomesticPlan();
                break;
            }
            case "COMMERCIAL" : {
                plan = new CommercialPlan();
                break;
            }
            case "INSTITUTIONAL" : {
                plan = new InstitutionalPlan();
                break;
            }
        }
        return plan;
    }
}
